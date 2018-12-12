package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/topherbullock/ci-example-go"
)

var _ = Describe("Main", func() {
	Describe("GetMessage", func() {
		It("returns the correct message", func() {
			Expect(GetMessage()).To(Equal("Hello go!"))
		})
	})
})
