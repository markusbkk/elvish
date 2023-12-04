// Command elvish is an alternative main program of Elvish that supports writing
// pprof profiles.
package main

import (
	"os"

	"github.com/markusbkk/elvish/pkg/buildinfo"
	"github.com/markusbkk/elvish/pkg/daemon"
	"github.com/markusbkk/elvish/pkg/lsp"
	"github.com/markusbkk/elvish/pkg/pprof"
	"github.com/markusbkk/elvish/pkg/prog"
	"github.com/markusbkk/elvish/pkg/shell"
)

func main() {
	os.Exit(prog.Run(
		[3]*os.File{os.Stdin, os.Stdout, os.Stderr}, os.Args,
		prog.Composite(
			&pprof.Program{}, &buildinfo.Program{}, &daemon.Program{}, &lsp.Program{},
			&shell.Program{ActivateDaemon: daemon.Activate})))
}
