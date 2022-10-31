package server_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"learntest/server"
)

var _ = Describe("Server", func() {

	var (
		body string
		addr string
	)

	Context("with no number passed", func() {
		BeforeEach(func() {
			addr = "http://localhost:8080/double"
			body = "value is missing"
		})

		It("should return 'value is missing' error", func() {
			req, err := http.NewRequest("GET", addr, nil)
			Expect(err).ShouldNot(HaveOccurred())

			rec := httptest.NewRecorder()
			server.DoubleHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			Expect(strings.TrimSpace(string(b))).To(Equal(body))
		})
	})

	Context("passing value 2", func() {
		BeforeEach(func() {
			addr = "http://localhost:8080/double?val=2"
			body = "4"
		})

		It("should return double of 2 which is 4", func() {
			req, err := http.NewRequest("GET", addr, nil)
			Expect(err).ShouldNot(HaveOccurred())

			rec := httptest.NewRecorder()
			server.DoubleHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.StatusCode).Should(Equal(http.StatusOK))
			Expect(strings.TrimSpace(string(b))).To(Equal(body))
		})
	})

	Context("passing a string value 'abc'", func() {
		BeforeEach(func() {
			addr = "http://localhost:8080/double?val=abc"
			body = "invalid number"
		})

		It("should return a string 'invalid number'", func() {
			req, err := http.NewRequest("GET", addr, nil)
			Expect(err).ShouldNot(HaveOccurred())

			rec := httptest.NewRecorder()
			server.DoubleHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			Expect(strings.TrimSpace(string(b))).To(Equal(body))
		})
	})

})
