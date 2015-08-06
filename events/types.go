package events

import (
	"net/http"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

type (
	//AppSearch - search for apps object
	EventSearch struct {
		Client           cloudFoundryClient
		ClientTargetInfo *cf.CloudFoundryAPIInfo
		EventsBlob       []cf.APIResponse
	}
	cloudFoundryClient interface {
		QueryAPIInfo() (*cf.CloudFoundryAPIInfo, error)
		QueryUsers(int, int, string, string) (userList cf.UserAPIResponse, err error)
		Query(verb string, domain string, path string, args interface{}) (response *http.Response)
	}
)
