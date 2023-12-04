// Command nav runs the navigation mode of the line editor.
package main

import (
	"fmt"

	"github.com/markusbkk/elvish/pkg/cli"
	"github.com/markusbkk/elvish/pkg/cli/modes"
	"github.com/markusbkk/elvish/pkg/cli/term"
	"github.com/markusbkk/elvish/pkg/cli/tk"
)

func main() {
	app := cli.NewApp(cli.AppSpec{})
	w, _ := modes.NewNavigation(app, modes.NavigationSpec{
		Bindings: tk.MapBindings{
			term.K('x'): func(tk.Widget) { app.CommitCode() },
		},
	})
	app.PushAddon(w)

	code, err := app.ReadCode()
	fmt.Println("code:", code)
	if err != nil {
		fmt.Println("err", err)
	}
}
