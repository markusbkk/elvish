package vals

import (
	"testing"

	. "github.com/markusbkk/elvish/pkg/tt"
)

func TestErrors(t *testing.T) {
	Test(t, Fn("error.Error", error.Error), Table{
		Args(cannotIterate{"num"}).Rets("cannot iterate num"),
		Args(cannotIterateKeysOf{"num"}).Rets("cannot iterate keys of num"),
	})
}
