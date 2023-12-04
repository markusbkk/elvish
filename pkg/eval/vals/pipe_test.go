package vals

import (
	"fmt"
	"testing"

	"github.com/markusbkk/elvish/pkg/persistent/hash"
	"github.com/markusbkk/elvish/pkg/testutil"
)

func TestPipe(t *testing.T) {
	pr, pw := testutil.MustPipe()
	defer pr.Close()
	defer pw.Close()

	TestValue(t, NewPipe(pr, pw)).
		Kind("pipe").
		Bool(true).
		Hash(hash.DJB(hash.UIntPtr(pr.Fd()), hash.UIntPtr(pw.Fd()))).
		Repr(fmt.Sprintf("<pipe{%v %v}>", pr.Fd(), pw.Fd())).
		Equal(NewPipe(pr, pw)).
		NotEqual(123, "a string", NewPipe(pw, pr)).
		AllKeys("r", "w").
		Index("r", pr).
		Index("w", pw)
}
