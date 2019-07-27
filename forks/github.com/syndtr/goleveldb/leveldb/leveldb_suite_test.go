package leveldb

import (
	"testing"

	"github.com/ie310mu/ie310go/forks/github.com/syndtr/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
