package ccclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	//URLPWSLogin - default pws login url
	URLPWSLogin = "https://login.run.pivotal.io"
	//RouteLogin - default oauth endpoint for cf
	RouteLogin     = "/oauth/token"
	dataUser       = "username"
	dataPass       = "password"
	dataScope      = "scope"
	dataScopeValue = ""
	dataGrant      = "grant_type"
	dataGrantValue = "password"
	HeaderAuth     = "Authorization"
)

//New - creates a new cloud controller client
func New(loginurl, user, pass string, client ClientDoer) *Client {
	return &Client{
		loginurl:     loginurl,
		user:         user,
		pass:         pass,
		client:       client,
		isStringData: false,
	}
}

type (
	ClientDoer interface {
		Do(*http.Request) (*http.Response, error)
	}
	//Client - cloud controller client object
	Client struct {
		Error        string `json:"error"`
		ErrorDesc    string `json:"error_description"`
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		Scope        string `json:"scope"`
		Jti          string `json:"jti"`
		user         string
		pass         string
		loginurl     string
		client       ClientDoer
		isStringData bool
	}
)

//Login - logs into the target cloud controller
func (s *Client) Login() (*Client, error) {
	var (
		err     error
		content []byte
	)

	if content, err = s.getToken(); err == nil {

		if json.Unmarshal(content, s); s.Error != "" {
			err = fmt.Errorf("error: %s - desc: %s", s.Error, s.ErrorDesc)
		}
	}
	return s, err
}

//CreateRequest - Creates a request object targeted at the cloud controller
func (s *Client) CreateRequest(verb, requestURL, path string, args interface{}) (*http.Request, error) {
	urlStr, dataBuf := s.createRequestData(requestURL, path, args)
	return http.NewRequest(verb, urlStr, dataBuf)
}

//HttpClient - returns the internal client object
func (s *Client) HttpClient() ClientDoer {
	return s.client
}

//CreateAuthRequest - Creates a request w/ auth token added to the header to allow authenticated calls to the cloud controller
func (s *Client) CreateAuthRequest(verb, requestURL, path string, args interface{}) (*http.Request, error) {
	req, err := s.CreateRequest(verb, requestURL, path, args)
	req.Header.Add(HeaderAuth, fmt.Sprintf("%s %s", s.TokenType, s.AccessToken))
	return req, err
}

func (s *Client) createRequestData(requestURL string, path string, postData interface{}) (apiUrl string, dataBuf *bytes.Buffer) {
	data := url.Values{}
	dataBuf = new(bytes.Buffer)

	if postData != nil {

		if d, ok := postData.(string); ok {
			dataBuf = bytes.NewBufferString(d)
		}

		if d, ok := postData.(map[string]string); ok {

			for i, v := range d {
				data.Add(i, v)
			}
			dataBuf = bytes.NewBufferString(data.Encode())
		}
	}
	u, _ := url.ParseRequestURI(requestURL)
	u.Path = path
	apiUrl = fmt.Sprintf("%v", u)
	return
}

func (s *Client) addLoginAuthHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add(HeaderAuth, "Basic Y2Y6")
}

func (s *Client) getToken() (content []byte, err error) {
	var res *http.Response
	var loginData = map[string]string{
		dataUser:  s.user,
		dataPass:  s.pass,
		dataScope: dataScopeValue,
		dataGrant: dataGrantValue,
	}
	req, _ := s.CreateRequest("POST", s.loginurl, RouteLogin, loginData)
	s.addLoginAuthHeaders(req)

	if res, err = s.client.Do(req); err == nil {
		defer res.Body.Close()
		content, err = ioutil.ReadAll(res.Body)
	}
	return
}
