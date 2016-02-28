package goapi_test

import (
	"crypto/tls"
	"github.com/geofffranks/goapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
)

func expect_body(req *http.Request, content string) {
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	Expect(err).ShouldNot(HaveOccurred())
	Expect(body).Should(Equal([]byte(content)))
}

var _ = Describe("HTTP Helpers", func() {
	Context("HttpRequest()", func() {
		It("should return an http.Request with encoded json if data provided", func() {
			req, err := goapi.HttpRequest("GET", "https://localhost:1234/test", map[string]interface{}{"asdf": 1234})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req).ShouldNot(BeNil())
			expect_body(req, `{"asdf":1234}`)
		})
		It("should return an http.Request without any payload if no data provided", func() {
			req, err := goapi.HttpRequest("GET", "https://localhost:1234/test", nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req).ShouldNot(BeNil())
			expect_body(req, "")
		})
		It("should fail if JSON data could not be marshaled", func() {
			req, err := goapi.HttpRequest("POST", "https://localhost:1234/test", map[int]string{1: "asdf"})
			Expect(err).Should(HaveOccurred())
			Expect(req).Should(BeNil())
		})
		It("should fail if http.NewRequest failed", func() {
			req, err := goapi.HttpRequest("INVALID", "%", nil) // '%' is an invalid URL!
			Expect(req).Should(BeNil())
			Expect(err).Should(HaveOccurred())
		})
		It("should set Content-Type + Accept headers", func() {
			req, err := goapi.HttpRequest("GET", "https://localhost:1234/test", nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req).ShouldNot(BeNil())

			Expect(req.Header.Get("Content-Type")).Should(Equal("application/json"))
			Expect(req.Header.Get("Accept")).Should(Equal("application/json"))
		})
		It("should use the specified method and URL in the request", func() {
			req, err := goapi.HttpRequest("GET", "https://localhost:1234/test", nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req).ShouldNot(BeNil())

			Expect(req.Method).Should(Equal("GET"))
			Expect(req.URL.String()).Should(Equal("https://localhost:1234/test"))

			req, err = goapi.HttpRequest("POST", "https://myhost:1234/stuff", "ping")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req).ShouldNot(BeNil())

			Expect(req.Method).Should(Equal("POST"))
			Expect(req.URL.String()).Should(Equal("https://myhost:1234/stuff"))
		})
	})
	Context("Get()", func() {
		It("should create a GET http.Request", func() {
			req, err := goapi.Get("https://localhost:1234/get")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req).ShouldNot(BeNil())
			Expect(req.Method).Should(Equal("GET"))
			Expect(req.URL.String()).Should(Equal("https://localhost:1234/get"))
			expect_body(req, "")
		})
		It("should return an error if unsuccessful", func() {
			req, err := goapi.Get("%")
			Expect(err).Should(HaveOccurred())
			Expect(req).Should(BeNil())
		})
	})
	Context("Post()", func() {
		It("should create a POST http.Request", func() {
			req, err := goapi.Post("https://localhost:1234/post", "teststring")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req).ShouldNot(BeNil())
			Expect(req.Method).Should(Equal("POST"))
			Expect(req.URL.String()).Should(Equal("https://localhost:1234/post"))
			expect_body(req, `"teststring"`)
		})
		It("should return an error if unsuccessful", func() {
			req, err := goapi.Post("%", nil)
			Expect(err).Should(HaveOccurred())
			Expect(req).Should(BeNil())
		})
	})
	Context("Put()", func() {
		It("should create a PUT http.Request", func() {
			req, err := goapi.Put("https://localhost:1234/put", "testput")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req.Method).Should(Equal("PUT"))
			Expect(req.URL.String()).Should(Equal("https://localhost:1234/put"))
			expect_body(req, `"testput"`)
		})
		It("should return an error if unsuccessful", func() {
			req, err := goapi.Put("%", nil)
			Expect(err).Should(HaveOccurred())
			Expect(req).Should(BeNil())
		})
	})
	Context("Patch()", func() {
		It("should create a PATCH http.Request", func() {
			req, err := goapi.Patch("https://localhost:1234/patch", "testpatch")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req).ShouldNot(BeNil())
			Expect(req.Method).Should(Equal("PATCH"))
			Expect(req.URL.String()).Should(Equal("https://localhost:1234/patch"))
			expect_body(req, `"testpatch"`)
		})
		It("should return an error if unsuccessful", func() {
			req, err := goapi.Patch("%", nil)
			Expect(err).Should(HaveOccurred())
			Expect(req).Should(BeNil())
		})
	})
	Context("Delete()", func() {
		It("should create a DELETE http.Request", func() {
			req, err := goapi.Delete("https://localhost:1234/delete")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(req).ShouldNot(BeNil())
			Expect(req.Method).Should(Equal("DELETE"))
			Expect(req.URL.String()).Should(Equal("https://localhost:1234/delete"))
			expect_body(req, "")
		})
		It("should return an error if unsuccessful", func() {
			req, err := goapi.Delete("%")
			Expect(err).Should(HaveOccurred())
			Expect(req).Should(BeNil())
		})
	})
	Context("Issue()", func() {
		It("should return errors from http.Client.Do()", func() {
			resp, err := goapi.Issue(&http.Request{})
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
		It("should return a parsed response upon success", func() {
			Skip("no tests written yet")
		})
	})
	Context("ParseResponse()", func() {
		It("should have tests", func() {
			Skip("No tests written yet")
		})
	})
	Context("Client()", func() {
		It("should return a generic client by default", func() {
			Expect(goapi.Client()).Should(Equal(&http.Client{}))
		})
		It("should allow setting a custom http.Client()", func() {
			client := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			}
			goapi.SetClient(client)

			Expect(goapi.Client()).Should(Equal(client))
		})
	})
	Context("BadResponseCode", func() {
		Context("Error()", func() {
			It("returns an error message", func() {
				err := goapi.BadResponseCode{
					StatusCode: 123,
					Message:    "this is an error",
					URL:        "https://localhost/test",
				}
				Expect(err.Error()).Should(Equal("https://localhost/test returned 123: this is an error"))

				err = goapi.BadResponseCode{
					StatusCode: 321,
					Message:    "this is a different error",
					URL:        "http://asdf.com/",
				}
				Expect(err.Error()).Should(Equal("http://asdf.com/ returned 321: this is a different error"))
			})
		})
	})
})
