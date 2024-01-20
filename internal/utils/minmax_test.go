package utils

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Min / Max", func() {
	It("returns the maximum", func() {
		Expect(Max(5, 7)).To(Equal(7))
		Expect(Max(5.5, 5.7)).To(Equal(5.7))
	})

	It("returns the minimum", func() {
		Expect(Min(5, 7)).To(Equal(5))
		Expect(Min(5.5, 5.7)).To(Equal(5.5))
	})
})
