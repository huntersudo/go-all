package gopher1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGopher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gopher Suite")
}
