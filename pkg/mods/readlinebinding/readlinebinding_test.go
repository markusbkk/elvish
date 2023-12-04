package readlinebinding_test

import (
	"os"
	"testing"

	"github.com/markusbkk/elvish/pkg/cli"
	"github.com/markusbkk/elvish/pkg/edit"
	"github.com/markusbkk/elvish/pkg/eval"
	. "github.com/markusbkk/elvish/pkg/eval/evaltest"
	"github.com/markusbkk/elvish/pkg/mods"
)

func TestReadlineBinding(t *testing.T) {
	// A smoke test to ensure that the readline-binding module has no errors.

	TestWithSetup(t, func(ev *eval.Evaler) {
		mods.AddTo(ev)
		ed := edit.NewEditor(cli.NewTTY(os.Stdin, os.Stderr), ev, nil)
		ev.ExtendBuiltin(eval.BuildNs().AddNs("edit", ed))
	},
		That("use readline-binding").DoesNothing(),
	)
}
