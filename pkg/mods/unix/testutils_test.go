//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package unix

import (
	"github.com/markusbkk/elvish/pkg/eval"
	"github.com/markusbkk/elvish/pkg/eval/evaltest"
)

var (
	That             = evaltest.That
	ErrorWithMessage = evaltest.ErrorWithMessage
)

func useUNIX(ev *eval.Evaler) {
	ev.ExtendGlobal(eval.BuildNs().AddNs("unix", Ns))
}
