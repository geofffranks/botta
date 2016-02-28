package goapi_test

import (
	"github.com/geofffranks/goapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Response Obj", func() {
	httpResponse := &http.Response{}
	response := goapi.Response{
		HTTPResponse: httpResponse,
		Raw:          []byte(`{"json":"content returned from server"}`),
		Data: map[string]interface{}{
			"string":  "asdf",
			"number":  1234,
			"boolean": true,
			"map": map[string]interface{}{
				"k": "v",
				"n": 1,
			},
			"array": []interface{}{
				1,
				2,
				"fdsa",
			},
		},
	}
	Context("Response", func() {
		It("should have publicly accessible Raw response", func() {
			Expect(response.Raw).Should(Equal([]byte(`{"json":"content returned from server"}`)))
		})
		It("should have publicly accessible pointer to the HTTP response", func() {
			Expect(response.HTTPResponse).Should(Equal(httpResponse))
		})
	})

	Context("Response.StringVal()", func() {
		It("should fail when specified path does not point to a string", func() {
			Skip("need working treefinder before tests will pass")
			str, err := response.StringVal("number")
			Expect(err).Should(HaveOccurred())
			Expect(str).Should(Equal(""))
		})
		It("should succeed when specified path points to a string", func() {
			Skip("need working treefinder before tests will pass")
			str, err := response.StringVal("string")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(str).Should(Equal("asdf"))
		})
	})

	Context("Response.NumVal()", func() {
		It("should fail when specified path does not point to a number", func() {
			Skip("need working treefinder before tests will pass")
			num, err := response.NumVal("string")
			Expect(err).Should(HaveOccurred())
			Expect(num).Should(Equal(0))
		})
		It("should succeed when specified path points to a number", func() {
			Skip("need working treefinder before tests will pass")
			num, err := response.NumVal("number")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(num).Should(Equal(1234))
		})
	})

	Context("Response.BoolVal()", func() {
		It("should fail when specified path does not point to a boolean", func() {
			Skip("need working treefinder before tests will pass")
			b, err := response.BoolVal("number")
			Expect(err).Should(HaveOccurred())
			Expect(b).Should(Equal(false))
		})
		It("should succeed when specified path points to a boolean", func() {
			Skip("need working treefinder before tests will pass")
			b, err := response.BoolVal("boolean")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(b).Should(BeTrue())
		})
	})

	Context("Response.MapVal()", func() {
		It("should fail when specified path does not point to a map", func() {
			Skip("need working treefinder before tests will pass")
			m, err := response.MapVal("number")
			Expect(err).Should(HaveOccurred())
			Expect(m).Should(Equal(nil))
		})
		It("should succeed when specified path points to a m", func() {
			Skip("need working treefinder before tests will pass")
			m, err := response.MapVal("map")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(m).Should(Equal(map[string]interface{}{
				"k": "v",
				"n": 1,
			}))
		})
	})

	Context("Response.ArrayVal()", func() {
		It("should fail when specified path does not point to an array", func() {
			Skip("need working treefinder before tests will pass")
			a, err := response.ArrayVal("number")
			Expect(err).Should(HaveOccurred())
			Expect(a).Should(Equal([]interface{}{}))
		})
		It("should succeed when specified path points to an array", func() {
			Skip("need working treefinder before tests will pass")
			a, err := response.ArrayVal("array")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(a).Should(Equal([]interface{}{
				1,
				2,
				"fdsa",
			}))
		})
	})

	Context("Response.Val()", func() {
		It("should fail when specified path does not point to something", func() {
			Skip("need working treefinder before tests will pass")
			i, err := response.Val("n'exist pas")
			Expect(err).Should(HaveOccurred())
			Expect(i).Should(Equal(nil))
		})
		It("should succeed when specified path points to something ", func() {
			Skip("need working treefinder before tests will pass")
			i, err := response.Val("string")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(i).Should(Equal(interface{}("asdf")))
		})
	})
})
