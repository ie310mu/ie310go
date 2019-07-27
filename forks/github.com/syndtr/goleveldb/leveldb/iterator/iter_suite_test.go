package iterator_test

import (
	"testing"

	"github.com/ie310mu/ie310go/forks/github.com/syndtr/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
