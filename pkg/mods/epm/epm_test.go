package epm_test

import (
	"testing"

	. "github.com/markusbkk/elvish/pkg/eval/evaltest"
	"github.com/markusbkk/elvish/pkg/mods"
)

func TestEPM(t *testing.T) {
	// A smoke test to ensure that the epm module has no errors.

	TestWithSetup(t, mods.AddTo,
		That("use epm").DoesNothing(),
	)
}
