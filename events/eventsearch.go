package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

//Init - initialize the state of your appsearch object
func (s *EventSearch) Init(client cloudFoundryClient) *EventSearch {
	s.Client = client
	s.ClientTargetInfo, _ = s.Client.QueryAPIInfo()
	return s
}

//CompileAllApps - compile the information for all of your applications
func (s *EventSearch) CompileRecentEvents() {
	var responseList cf.APIResponseList
	queryArgs := fmt.Sprintf("order-direction=desc&results-per-page=%d", eventResultsLimit)
	res := s.Client.Query("GET", s.ClientTargetInfo.APIEndpoint, eventRESTPath, queryArgs)
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyBytes, &responseList)
	s.EventsBlob = responseList.Resources
}
