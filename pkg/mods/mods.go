// Package mods collects standard library modules.
package mods

import (
	"github.com/markusbkk/elvish/pkg/eval"
	"github.com/markusbkk/elvish/pkg/mods/epm"
	"github.com/markusbkk/elvish/pkg/mods/file"
	"github.com/markusbkk/elvish/pkg/mods/flag"
	"github.com/markusbkk/elvish/pkg/mods/math"
	"github.com/markusbkk/elvish/pkg/mods/path"
	"github.com/markusbkk/elvish/pkg/mods/platform"
	"github.com/markusbkk/elvish/pkg/mods/re"
	"github.com/markusbkk/elvish/pkg/mods/readlinebinding"
	"github.com/markusbkk/elvish/pkg/mods/str"
	"github.com/markusbkk/elvish/pkg/mods/unix"
)

// AddTo adds all standard library modules to the Evaler.
func AddTo(ev *eval.Evaler) {
	ev.AddModule("math", math.Ns)
	ev.AddModule("path", path.Ns)
	ev.AddModule("platform", platform.Ns)
	ev.AddModule("re", re.Ns)
	ev.AddModule("str", str.Ns)
	ev.AddModule("file", file.Ns)
	ev.AddModule("flag", flag.Ns)
	if unix.ExposeUnixNs {
		ev.AddModule("unix", unix.Ns)
	}
	ev.BundledModules["epm"] = epm.Code
	ev.BundledModules["readline-binding"] = readlinebinding.Code
}
