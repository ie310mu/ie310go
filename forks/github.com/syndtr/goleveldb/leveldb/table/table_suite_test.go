package table

import (
	"testing"

	"github.com/ie310mu/ie310go/forks/github.com/syndtr/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
