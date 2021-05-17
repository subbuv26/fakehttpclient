package fakehttpclient

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFakeHTTPClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fake HTTP Client Suite")
}

var _= Describe("Fake HTTP Client Tests", func() {
	Describe("Getting New Client", func() {
		var responses []*http.Response
		BeforeEach(func() {
			responses = []*http.Response{
				{
					StatusCode: 200,
					Header:     http.Header{},
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("body"))),
				},
			}
		})

		It("Create a client", func() {

			client, err := NewFakeHTTPClient(responses, 2)
			Expect(err).To(BeNil(), "Failed to Create Client")
			Expect(client).NotTo(BeNil(), "Failed to Create Client")
		})

		It("Empty Responses", func() {
			client, err := NewFakeHTTPClient(nil, 2)
			Expect(err).NotTo(BeNil(), "Failed to Validate responses")
			Expect(client).To(BeNil(), "Failed to Validate responses")
		})

		It("Invalid maxRun (0)", func() {
			client, err := NewFakeHTTPClient(responses, 0)
			Expect(err).NotTo(BeNil(), "Failed to Validate responses")
			Expect(client).To(BeNil(), "Failed to Validate responses")
		})

		It("Invalid maxRun (-ve)", func() {
			client, err := NewFakeHTTPClient(responses, -2)
			Expect(err).NotTo(BeNil(), "Failed to Validate responses")
			Expect(client).To(BeNil(), "Failed to Validate responses")
		})
	})

	Describe("Send Requests with the Client", func() {
		var client *http.Client
		var resp1, resp2, resp3 *http.Response
		var req *http.Request

		BeforeEach(func() {
			resp1 = &http.Response{
					StatusCode: 200,
					Header:     http.Header{},
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("body"))),
				}

			resp2 = &http.Response{
					StatusCode: 404,
					Header:     http.Header{},
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("body"))),
				}

			resp3 = &http.Response{
					StatusCode: 505,
					Header:     http.Header{},
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("body"))),
				}

			req, _ = http.NewRequest("POST", "http://1.2.3.4",  bytes.NewBuffer([]byte("{}")))
		})

		It("Basic Request and Response", func() {
			responses := []*http.Response{resp1}
			client, _ = NewFakeHTTPClient(responses, 1)
			resp, err := client.Do(req)
			Expect(err).To(BeNil())
			Expect(resp).To(Equal(resp1))
		})

		It("Continuous Requests and Responses", func() {
			responses := []*http.Response{resp1, resp2, resp3}
			client, _ = NewFakeHTTPClient(responses, 3)
			for _, resp := range responses {
				rsp, err := client.Do(req)
				Expect(err).To(BeNil())
				Expect(rsp).To(Equal(resp))
			}
		})

		It("Run over multiple Cycles", func() {
			responses := []*http.Response{resp1, resp2, resp3}
			client, _ = NewFakeHTTPClient(responses, 9)

			for i := 0; i < 3; i++ {
				for _, resp := range responses {
					rsp, err := client.Do(req)
					Expect(err).To(BeNil())
					Expect(rsp).To(Equal(resp))
				}
			}
		})

		It("Out Run maxRun", func() {
			responses := []*http.Response{resp1, resp2, resp3}
			client, _ = NewFakeHTTPClient(responses, 3)
			for _, resp := range responses {
				rsp, err := client.Do(req)
				Expect(err).To(BeNil())
				Expect(rsp).To(Equal(resp))
			}
			rsp, err := client.Do(req)
			Expect(err).NotTo(BeNil())
			Expect(rsp).To(BeNil())
		})

		It("Nil Response", func() {
			responses := []*http.Response{nil}
			client, _ = NewFakeHTTPClient(responses, 1)
			resp, err := client.Do(req)
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})
	})
})
