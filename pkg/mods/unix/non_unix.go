//go:build windows || plan9 || js
// +build windows plan9 js

// Package unix exports an Elvish namespace that contains variables and
// functions that deal with features unique to UNIX-like operating systems. On
// non-UNIX operating systems it exports an empty namespace.
package unix

import (
	"github.com/markusbkk/elvish/pkg/eval"
)

// ExposeUnixNs indicate whether this module should be exposed as a usable
// elvish namespace.
const ExposeUnixNs = false

// Ns is an Elvish namespace that contains variables and functions that deal
// with features unique to UNIX-like operating systems. On
var Ns = &eval.Ns{}
