package proto_test

import (
	"testing"

	. "github.com/ie310mu/ie310go/forks/github.com/onsi/ginkgo"
	. "github.com/ie310mu/ie310go/forks/github.com/onsi/gomega"
)

func TestGinkgoSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "proto")
}
