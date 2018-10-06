package bowling_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBowling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bowling Suite")
}
