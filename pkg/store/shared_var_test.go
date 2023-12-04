package store_test

import (
	"testing"

	"github.com/markusbkk/elvish/pkg/store"
	"github.com/markusbkk/elvish/pkg/store/storetest"
)

func TestSharedVar(t *testing.T) {
	storetest.TestSharedVar(t, store.MustTempStore(t))
}
