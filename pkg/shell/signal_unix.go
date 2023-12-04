//go:build !windows && !plan9
// +build !windows,!plan9

package shell

import (
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/markusbkk/elvish/pkg/sys"
)

func handleSignal(sig os.Signal, stderr io.Writer) {
	switch sig {
	case syscall.SIGHUP:
		syscall.Kill(0, syscall.SIGHUP)
		os.Exit(0)
	case syscall.SIGUSR1:
		fmt.Fprint(stderr, sys.DumpStack())
	}
}
