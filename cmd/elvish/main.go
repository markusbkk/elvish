// Elvish is a cross-platform shell, supporting Linux, BSDs and Windows. It
// features an expressive programming language, with features like namespacing
// and anonymous functions, and a fully programmable user interface with
// friendly defaults. It is suitable for both interactive use and scripting.
package elvish

import (
	"os"

	"github.com/markusbkk/elvish/pkg/buildinfo"
	"github.com/markusbkk/elvish/pkg/daemon"
	"github.com/markusbkk/elvish/pkg/lsp"
	"github.com/markusbkk/elvish/pkg/prog"
	"github.com/markusbkk/elvish/pkg/shell"
	"github.com/icexin/eggos/app"
)

func elvish(ctx *app.Context) error {
	err := ctx.ParseFlags()
	if err != nil {
		return err
	}
	os.Exit(prog.Run(
		[3]*os.File{os.Stdin, os.Stdout, os.Stderr}, os.Args,
		prog.Composite(
			&buildinfo.Program{}, &daemon.Program{}, &lsp.Program{},
			&shell.Program{ActivateDaemon: daemon.Activate})))
	return nil
}

func init() {
	app.Register("elvish", elvish)
}
