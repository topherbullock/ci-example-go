package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCiExampleGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CiExampleGo Suite")
}
