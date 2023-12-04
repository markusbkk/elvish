package parse_test

import (
	"testing"

	"github.com/markusbkk/elvish/pkg/eval/vals"
	. "github.com/markusbkk/elvish/pkg/parse"
)

func TestSourceAsStructMap(t *testing.T) {
	vals.TestValue(t, Source{Name: "[tty]", Code: "echo"}).
		Kind("structmap").
		Repr("[&name='[tty]' &code=<...> &is-file=$false]").
		AllKeys("name", "code", "is-file")

	vals.TestValue(t, Source{Name: "/etc/rc.elv", Code: "echo", IsFile: true}).
		Index("is-file", true)
}
