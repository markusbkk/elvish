package clitest

import (
	"testing"

	"github.com/markusbkk/elvish/pkg/cli"
	"github.com/markusbkk/elvish/pkg/cli/term"
	"github.com/markusbkk/elvish/pkg/cli/tk"
	"github.com/markusbkk/elvish/pkg/ui"
)

func TestFixture(t *testing.T) {
	f := Setup(
		WithSpec(func(spec *cli.AppSpec) {
			spec.CodeAreaState.Buffer = tk.CodeBuffer{Content: "test", Dot: 4}
		}),
		WithTTY(func(tty TTYCtrl) {
			tty.SetSize(20, 30) // h = 20, w = 30
		}),
	)
	defer f.Stop()

	// Verify that the functions passed to Setup have taken effect.
	if f.App.ActiveWidget().(tk.CodeArea).CopyState().Buffer.Content != "test" {
		t.Errorf("WithSpec did not work")
	}

	buf := f.MakeBuffer()
	// Verify that the WithTTY function has taken effect.
	if buf.Width != 30 {
		t.Errorf("WithTTY did not work")
	}

	f.TestTTY(t, "test", term.DotHere)

	f.App.Notify(ui.T("something"))
	f.TestTTYNotes(t, "something")

	f.App.CommitCode()
	if code, err := f.Wait(); code != "test" || err != nil {
		t.Errorf("Wait returned %q, %v", code, err)
	}
}
