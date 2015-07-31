package applications

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

func (s *AppSearch) Init(client cloudFoundryClient) *AppSearch {
	s.Client = client
	s.ClientTargetInfo, _ = s.Client.QueryAPIInfo()
	s.AppStats = new(AppAggregate)
	s.GetAppCount()
	return s
}

func (s *AppSearch) GetAppCount() (appCount int) {
	var responseList cf.APIResponseList
	res := s.Client.Query("GET", s.ClientTargetInfo.APIEndpoint, appRESTPath, nil)
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyBytes, &responseList)
	s.AppStats.TotalAppCount = responseList.TotalResults
	return s.AppStats.TotalAppCount
}

func (s *AppSearch) CompileAllApps() {
	var responseList cf.APIResponseList
	res := s.Client.Query("GET", s.ClientTargetInfo.APIEndpoint, appRESTPath, fmt.Sprintf("results-per-page=%d", s.AppStats.TotalAppCount))
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyBytes, &responseList)

	for _, applicationRecord := range responseList.Resources {
		s.processApplicationRecord(applicationRecord)
	}
}

func (s *AppSearch) processApplicationRecord(appRecord cf.APIResponse) {
	s.processAI(appRecord)
	s.processBP(appRecord)
}

func (s *AppSearch) processAI(appRecord cf.APIResponse) {
	s.AppStats.TotalInstanceCount += int(appRecord.Entity["instances"].(float64))
}

func (s *AppSearch) processBP(appRecord cf.APIResponse) {
	var buildpackName string

	if appRecord.Entity[buildpackFieldname] != nil {
		buildpackName = appRecord.Entity[buildpackFieldname].(string)
	}
	s.checkAgainstRegisteredBuildpacks(buildpackName)
}

func (s *AppSearch) checkAgainstRegisteredBuildpacks(buildpackName string) {
	if strings.Contains(buildpackName, javaMatcher) {
		s.AppStats.JavaBPCount++

	} else if strings.Contains(buildpackName, rubyMatcher) {
		s.AppStats.RubyBPCount++

	} else if strings.Contains(buildpackName, nodeMatcher) {
		s.AppStats.NodeBPCount++

	} else if strings.Contains(buildpackName, goMatcher) {
		s.AppStats.GOBPCount++

	} else if strings.Contains(buildpackName, pyMatcher) {
		s.AppStats.PythonBPCount++

	} else if strings.Contains(buildpackName, phpMatcher) {
		s.AppStats.PHPBPCount++

	} else {
		s.AppStats.OtherBPCount++
	}
}
