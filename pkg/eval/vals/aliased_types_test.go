package vals

import (
	"testing"

	"github.com/markusbkk/elvish/pkg/testutil"
	. "github.com/markusbkk/elvish/pkg/tt"
)

func TestMakeMap_PanicsWithOddNumberOfArguments(t *testing.T) {
	Test(t, Fn("Recover", testutil.Recover), Table{
		//lint:ignore SA5012 testing panic
		Args(func() { MakeMap("foo") }).Rets("odd number of arguments to MakeMap"),
	})
}
