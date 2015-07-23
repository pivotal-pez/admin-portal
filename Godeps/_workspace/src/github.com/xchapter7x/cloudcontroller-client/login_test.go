package ccclient_test

import (
	"bytes"
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/cloudcontroller-client"
)

var _ = Describe("New cloud controller client", func() {
	var (
		sampleSuccessTokenStringFormat = `{"access_token":"%s","token_type":"bearer","refresh_token":"%s","expires_in":599,"scope":"password.write cloud_controller.write openid cloud_controller.read","jti":"%s"}`
		sampleFailureTokenString       = `{"error":"unauthorized","error_description":"Authentication failed"}`
		sampleUser                     = "randomeuser@pivotal.io"
		samplePass                     = "mypass"
	)

	Context("when calling login w/ valid arguments", func() {
		var (
			client              *Client
			err                 error
			controlAccessToken  = "my control access token"
			controlRefreshToken = "my control refresh token"
			controlJti          = "my control jti"
		)

		BeforeEach(func() {
			client, err = New(URLPWSLogin, sampleUser, samplePass, &mockDoer{
				res: &http.Response{
					Body: nopCloser{bytes.NewBufferString(fmt.Sprintf(sampleSuccessTokenStringFormat, controlAccessToken, controlRefreshToken, controlJti))},
				},
			}).Login()
		})

		It("should return a nil error", func() {
			Ω(err).Should(BeNil())
		})

		It("should set proper values (AccessToken) from the response token object", func() {
			Ω(client.AccessToken).Should(Equal(controlAccessToken))
		})
	})

	Context("when calling login w/ in-valid arguments", func() {
		var (
			client *Client
			err    error
		)

		BeforeEach(func() {
			client, err = New(URLPWSLogin, sampleUser, samplePass, &mockDoer{
				res: &http.Response{
					Body: nopCloser{bytes.NewBufferString(sampleFailureTokenString)},
				},
			}).Login()
		})

		It("should return a NON-nil error", func() {
			Ω(err).ShouldNot(BeNil())
		})

		It("should not set token", func() {
			Ω(client.AccessToken).Should(BeEmpty())
		})
	})

	XContext("non-functional integration test", func() {
		It("should not run right now", func() {
			client := New(URLPWSLogin, "jcalabrese@pivotal.io", "pass", new(http.Client))
			fmt.Println(client.Login())
			Ω(true).Should(BeFalse())
		})
	})
})
