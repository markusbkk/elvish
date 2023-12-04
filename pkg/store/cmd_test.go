package store_test

import (
	"testing"

	"github.com/markusbkk/elvish/pkg/store"
	"github.com/markusbkk/elvish/pkg/store/storetest"
)

func TestCmd(t *testing.T) {
	storetest.TestCmd(t, store.MustTempStore(t))
}
