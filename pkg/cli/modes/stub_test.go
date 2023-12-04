package modes

import (
	"testing"
	"time"

	"github.com/markusbkk/elvish/pkg/cli"
	. "github.com/markusbkk/elvish/pkg/cli/clitest"
	"github.com/markusbkk/elvish/pkg/cli/term"
	"github.com/markusbkk/elvish/pkg/cli/tk"
)

func TestStub_Rendering(t *testing.T) {
	f := Setup()
	defer f.Stop()

	startStub(f.App, StubSpec{Name: " STUB "})
	f.TestTTY(t,
		"", term.DotHere, "\n",
		" STUB ", Styles,
		"******",
	)
}

func TestStub_Handling(t *testing.T) {
	f := Setup()
	defer f.Stop()

	bindingCalled := make(chan bool)
	startStub(f.App, StubSpec{
		Bindings: tk.MapBindings{
			term.K('a'): func(tk.Widget) { bindingCalled <- true }},
	})

	f.TTY.Inject(term.K('a'))
	select {
	case <-bindingCalled:
		// OK
	case <-time.After(time.Second):
		t.Errorf("Handler not called after 1s")
	}
}

func startStub(app cli.App, spec StubSpec) {
	w := NewStub(spec)
	app.PushAddon(w)
	app.Redraw()
}
