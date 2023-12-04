package eval_test

import (
	"testing"

	. "github.com/markusbkk/elvish/pkg/eval"
	. "github.com/markusbkk/elvish/pkg/eval/evaltest"
	"github.com/markusbkk/elvish/pkg/eval/vals"
)

func TestExplicitBuiltinModule(t *testing.T) {
	TestWithSetup(t, func(ev *Evaler) { ev.Args = vals.MakeList("a", "b") },
		That("all $args").Puts("a", "b"),
		// Regression test for #1414
		That("use builtin; all $builtin:args").Puts("a", "b"),
	)
}
