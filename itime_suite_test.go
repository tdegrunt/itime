package itime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestItime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Itime Suite")
}
