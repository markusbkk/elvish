package vars

// Note: This doesn't have an associated env_list_tests.go because most of its functionality is
// tested by TestSetEnv_PATH and related tests.

import (
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/markusbkk/elvish/pkg/diag"
	"github.com/markusbkk/elvish/pkg/eval/vals"
)

var (
	pathListSeparator = string(os.PathListSeparator)
	forbiddenInPath   = pathListSeparator + "\x00"
)

// Errors
var (
	ErrPathMustBeString          = errors.New("path must be string")
	ErrPathContainsForbiddenChar = errors.New("path cannot contain NUL byte, colon on UNIX or semicolon on Windows")
)

// NewEnvListVar returns a variable whose value is a list synchronized with an
// environment variable with the elements joined by os.PathListSeparator.
//
// Elements in the value of the variable must be strings, and cannot contain
// os.PathListSeparator or \0; attempting to put any in its elements will result in
// an error.
func NewEnvListVar(name string) Var {
	return &envListVar{envName: name}
}

type envListVar struct {
	sync.RWMutex
	envName    string
	cacheFor   string
	cacheValue interface{}
}

// Get returns a Value for an EnvPathList.
func (envli *envListVar) Get() interface{} {
	envli.Lock()
	defer envli.Unlock()

	value := os.Getenv(envli.envName)
	if value == envli.cacheFor {
		return envli.cacheValue
	}
	envli.cacheFor = value
	v := vals.EmptyList
	for _, path := range strings.Split(value, pathListSeparator) {
		v = v.Conj(path)
	}
	envli.cacheValue = v
	return envli.cacheValue
}

// Set sets an EnvPathList. The underlying environment variable is set.
func (envli *envListVar) Set(v interface{}) error {
	var (
		paths      []string
		errElement error
	)
	errIterate := vals.Iterate(v, func(v interface{}) bool {
		s, ok := v.(string)
		if !ok {
			errElement = ErrPathMustBeString
			return false
		}
		path := s
		if strings.ContainsAny(path, forbiddenInPath) {
			errElement = ErrPathContainsForbiddenChar
			return false
		}
		paths = append(paths, s)
		return true
	})

	if errElement != nil || errIterate != nil {
		return diag.Errors(errElement, errIterate)
	}

	envli.Lock()
	defer envli.Unlock()
	os.Setenv(envli.envName, strings.Join(paths, pathListSeparator))
	return nil
}
