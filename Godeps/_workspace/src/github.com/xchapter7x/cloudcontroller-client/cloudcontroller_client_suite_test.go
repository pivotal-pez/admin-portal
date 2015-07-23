package ccclient_test

import (
	"io"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCloudControllerClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cloud Controller Client Suite")
}

type mockDoer struct {
	req *http.Request
	res *http.Response
	err error
}

func (s *mockDoer) Do(req *http.Request) (*http.Response, error) {
	return s.res, s.err
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
