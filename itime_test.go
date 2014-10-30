package itime_test

import (
	. "github.com/tdegrunt/itime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Itime", func() {
	Describe("Instantiation", func() {
		It("should instantiate from hours and minutes", func() {
			now := New(23, 31)
			hours, minutes := now.Split()
			Expect(hours).To(Equal(23))
			Expect(minutes).To(Equal(31))
			Expect(now.Int()).To(Equal(2331))
		})
		It("should instantiate from integer", func() {
			now := Itime(2331)
			hours, minutes := now.Split()
			Expect(hours).To(Equal(23))
			Expect(minutes).To(Equal(31))
			Expect(now.Int()).To(Equal(2331))
		})
	})
	Describe("Calculation", func() {
		It("should subtract times", func() {
			subbed := Itime(2311).Sub(Itime(1330))
			hours, minutes := subbed.Split()
			Expect(hours).To(Equal(9))
			Expect(minutes).To(Equal(41))
			Expect(subbed.Int()).To(Equal(941))
		})
		It("should add times", func() {
			added := Itime(2341).Add(Itime(1350))
			hours, minutes := added.Split()
			Expect(hours).To(Equal(37))
			Expect(minutes).To(Equal(31))
			Expect(added.Int()).To(Equal(3731))
		})
	})
	Describe("Comparison", func() {
		It("should report that 1330 is before 2311", func() {
			before := Itime(1330).Before(Itime(2311))
			Expect(before).To(BeTrue())
		})
		It("should report that 2311 is not before 1330", func() {
			before := Itime(2311).Before(Itime(1330))
			Expect(before).To(BeFalse())
		})
		It("should report that 2311 is after 1330", func() {
			after := Itime(2311).After(Itime(1330))
			Expect(after).To(BeTrue())
		})
		It("should report that 1330 is not after 2311", func() {
			after := Itime(1330).After(Itime(2311))
			Expect(after).To(BeFalse())
		})

		It("should report that 2311 is equal to 2311", func() {
			equal := Itime(2311).Equal(Itime(2311))
			Expect(equal).To(BeTrue())
		})
		It("should report that 1330 is not equal 2311", func() {
			equal := Itime(1330).Equal(Itime(2311))
			Expect(equal).To(BeFalse())
		})
	})
	Describe("Duration", func() {
		It("should return a duration of 13h30m0s for 1330", func() {
			duration := Itime(1330).Duration()
			Expect(duration.String()).To(Equal("13h30m0s"))
		})

	})
})
