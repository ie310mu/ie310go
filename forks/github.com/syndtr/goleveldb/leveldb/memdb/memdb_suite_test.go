package memdb

import (
	"testing"

	"github.com/ie310mu/ie310go/forks/github.com/syndtr/goleveldb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
