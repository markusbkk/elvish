// Command elvish is an alternative main program of Elvish that does not include
// the daemon subprogram.
package main

import (
	"os"

	"github.com/markusbkk/elvish/pkg/buildinfo"
	"github.com/markusbkk/elvish/pkg/lsp"
	"github.com/markusbkk/elvish/pkg/prog"
	"github.com/markusbkk/elvish/pkg/shell"
)

func main() {
	os.Exit(prog.Run(
		[3]*os.File{os.Stdin, os.Stdout, os.Stderr}, os.Args,
		prog.Composite(&buildinfo.Program{}, &lsp.Program{}, &shell.Program{})))
}
