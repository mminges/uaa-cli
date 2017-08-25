package utils_test

import (

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jhamon/uaa-cli/utils"
)

var _ = Describe("UrlHelpers", func() {
	Describe("BuildUrl", func() {
		It("adds path to base url", func() {
			url, _ := utils.BuildUrl("http://localhost:8080", "foo")
			Expect(url.String()).To(Equal("http://localhost:8080/foo"))

			url, _ = utils.BuildUrl("http://localhost:8080/", "foo")
			Expect(url.String()).To(Equal("http://localhost:8080/foo"))

			url, _ = utils.BuildUrl("http://localhost:8080/", "/foo")
			Expect(url.String()).To(Equal("http://localhost:8080/foo"))

			url, _ = utils.BuildUrl("http://localhost:8080", "/foo")
			Expect(url.String()).To(Equal("http://localhost:8080/foo"))
		})
	})
})