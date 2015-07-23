package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

//ListUserOrg - lists the matches from a query on guid for orgs
func (s *UserSearch) ListUserOrgs(userGUID string) (resObj *cf.APIResponseList, err error) {
	path := fmt.Sprintf("/v2/users/%s/organizations", userGUID)

	if res := s.Client.Query("GET", s.ClientTargetInfo.APIEndpoint, path, nil); res.StatusCode == OrgsSuccessStatusCode {
		b, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(b, &resObj)

	} else {
		err = ErrListUserOrgs
	}
	return
}

//ListUserSpaces - lists the matches from a query on guid for spaces
func (s *UserSearch) ListUserSpaces(userGUID string) (resObj *cf.APIResponseList, err error) {
	path := fmt.Sprintf("/v2/users/%s/spaces", userGUID)

	if res := s.Client.Query("GET", s.ClientTargetInfo.APIEndpoint, path, nil); res.StatusCode == SpacesSuccessStatusCode {
		b, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(b, &resObj)

	} else {
		err = ErrListUserSpaces
	}
	return
}
