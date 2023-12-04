package pprof_test

import (
	"os"
	"testing"

	. "github.com/markusbkk/elvish/pkg/pprof"
	"github.com/markusbkk/elvish/pkg/prog"
	"github.com/markusbkk/elvish/pkg/prog/progtest"
	"github.com/markusbkk/elvish/pkg/testutil"
)

var (
	Test       = progtest.Test
	ThatElvish = progtest.ThatElvish
)

func TestProgram(t *testing.T) {
	testutil.InTempDir(t)

	Test(t, prog.Composite(&Program{}, noopProgram{}),
		ThatElvish("-cpuprofile", "cpuprof").DoesNothing(),
		ThatElvish("-cpuprofile", "/a/bad/path").
			WritesStderrContaining("Warning: cannot create CPU profile:"),
	)

	// Check for the effect of -cpuprofile. There isn't much to test beyond a
	// sanity check that the profile file now exists.
	_, err := os.Stat("cpuprof")
	if err != nil {
		t.Errorf("CPU profile file does not exist: %v", err)
	}
}

type noopProgram struct{}

func (noopProgram) RegisterFlags(*prog.FlagSet)     {}
func (noopProgram) Run([3]*os.File, []string) error { return nil }
