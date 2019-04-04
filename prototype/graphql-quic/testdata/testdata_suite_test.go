package testdata

// TAKEN FROM: https://github.com/lucas-clemente/quic-go/tree/master/internal/testdata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTestdata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testdata Suite")
}
