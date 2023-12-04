package vars

import (
	"testing"

	"github.com/markusbkk/elvish/pkg/eval/errs"
	"github.com/markusbkk/elvish/pkg/tt"
)

var Args = tt.Args

func TestNewReadOnly(t *testing.T) {
	v := NewReadOnly("haha")
	if v.Get() != "haha" {
		t.Errorf("Get doesn't return initial value")
	}

	err := v.Set("lala")
	if _, ok := err.(errs.SetReadOnlyVar); !ok {
		t.Errorf("Set a readonly var doesn't error as expected: %#v", err)
	}
}

func TestIsReadOnly(t *testing.T) {
	tt.Test(t, tt.Fn("IsReadOnly", IsReadOnly), tt.Table{
		Args(NewReadOnly("foo")).Rets(true),
		Args(FromGet(func() interface{} { return "foo" })).Rets(true),
		Args(FromInit("foo")).Rets(false),
	})
}
