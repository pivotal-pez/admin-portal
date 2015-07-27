package users

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

type (
	//UserSearch - a user search object
	UserSearch struct {
		Client           cloudFoundryClient
		ClientTargetInfo *cf.CloudFoundryAPIInfo
	}
	cloudFoundryClient interface {
		QueryAPIInfo() (*cf.CloudFoundryAPIInfo, error)
		QueryUsers(int, int, string, string) (userList cf.UserAPIResponse, err error)
		Query(verb string, domain string, path string, args interface{}) (response *http.Response)
	}
)

//Init - initialize the user search object
func (s *UserSearch) Init(client cloudFoundryClient) *UserSearch {
	s.Client = client
	s.ClientTargetInfo, _ = s.Client.QueryAPIInfo()
	return s
}

//BuildQuery - construct a query string
func (s *UserSearch) BuildQuery(usertype, username string) (query string) {
	l := []string{}

	if usertype != "" {
		l = append(l, fmt.Sprintf("origin eq '%s'", usertype))
	}

	if username != "" {
		l = append(l, fmt.Sprintf("userName co '%s'", username))
	}
	query = url.QueryEscape(strings.Join(l, " and "))
	return
}

//List - lists the matches from a query on usertype/username
func (s *UserSearch) List(usertype, username string) (users cf.UserAPIResponse, err error) {
	var u cf.UserAPIResponse

	if u, err = s.Client.QueryUsers(1, 1, "id", ""); err == nil {
		query := s.BuildQuery(usertype, username)
		users, err = s.Client.QueryUsers(1, u.TotalResults, "", query)
	}
	return
}
