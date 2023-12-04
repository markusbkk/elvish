package store_test

import (
	"testing"

	"github.com/markusbkk/elvish/pkg/store"
	"github.com/markusbkk/elvish/pkg/store/storetest"
)

func TestDir(t *testing.T) {
	storetest.TestDir(t, store.MustTempStore(t))
}
