package edit

import (
	"testing"

	"github.com/markusbkk/elvish/pkg/cli/term"
	"github.com/markusbkk/elvish/pkg/store/storedefs"
	"github.com/markusbkk/elvish/pkg/ui"
)

func TestHistWalk_Up_EndOfHistory(t *testing.T) {
	f := startHistwalkTest(t)

	f.TTYCtrl.Inject(term.K(ui.Up))
	f.TestTTYNotes(t,
		"error: end of history", Styles,
		"!!!!!!")
}

func TestHistWalk_Down_EndOfHistory(t *testing.T) {
	f := startHistwalkTest(t)

	// Not bound by default, so we need to use evals.
	evals(f.Evaler, `edit:history:down`)
	f.TestTTYNotes(t,
		"error: end of history", Styles,
		"!!!!!!")
}

func TestHistWalk_Accept(t *testing.T) {
	f := startHistwalkTest(t)

	f.TTYCtrl.Inject(term.K(ui.Right))
	f.TestTTY(t,
		"~> echo a", Styles,
		"   vvvv  ", term.DotHere,
	)
}

func TestHistWalk_Close(t *testing.T) {
	f := startHistwalkTest(t)

	f.TTYCtrl.Inject(term.K('[', ui.Ctrl))
	f.TestTTY(t, "~> ", term.DotHere)
}

func TestHistWalk_DownOrQuit(t *testing.T) {
	f := startHistwalkTest(t)

	f.TTYCtrl.Inject(term.K(ui.Down))
	f.TestTTY(t, "~> ", term.DotHere)
}

func TestHistory_FastForward(t *testing.T) {
	f := setup(t, storeOp(func(s storedefs.Store) {
		s.AddCmd("echo a")
	}))

	f.Store.AddCmd("echo b")
	evals(f.Evaler, `edit:history:fast-forward`)
	f.TTYCtrl.Inject(term.K(ui.Up))
	f.TestTTY(t,
		"~> echo b", Styles,
		"   VVVV__", term.DotHere, "\n",
		" HISTORY #2 ", Styles,
		"************",
	)
}

func startHistwalkTest(t *testing.T) *fixture {
	// The part of the test shared by all tests.
	f := setup(t, storeOp(func(s storedefs.Store) {
		s.AddCmd("echo a")
	}))

	f.TTYCtrl.Inject(term.K(ui.Up))
	f.TestTTY(t,
		"~> echo a", Styles,
		"   VVVV__", term.DotHere, "\n",
		" HISTORY #1 ", Styles,
		"************",
	)
	return f
}
