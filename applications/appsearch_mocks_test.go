package applications_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

type fakeCloudFoundryClient struct {
}

func (s *fakeCloudFoundryClient) QueryAPIInfo() (apiInfo *cf.CloudFoundryAPIInfo, err error) {
	return new(cf.CloudFoundryAPIInfo), nil
}

func (s *fakeCloudFoundryClient) QueryUsers(int, int, string, string) (userList cf.UserAPIResponse, err error) {
	return
}

func (s *fakeCloudFoundryClient) Query(verb string, domain string, path string, args interface{}) (response *http.Response) {
	fileBytes, _ := ioutil.ReadFile("fixtures/applications.json")
	response = &http.Response{
		Body: nopCloser{bytes.NewBuffer(fileBytes)},
	}
	return
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
