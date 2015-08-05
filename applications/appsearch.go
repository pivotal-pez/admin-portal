package applications

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

//Init - initialize the state of your appsearch object
func (s *AppSearch) Init(client cloudFoundryClient) *AppSearch {
	s.Client = client
	s.ClientTargetInfo, _ = s.Client.QueryAPIInfo()
	s.AppStats = new(AppAggregate)
	s.GetAppCount()
	return s
}

//GetAppCount - get the current application count
func (s *AppSearch) GetAppCount() (appCount int) {
	var responseList cf.APIResponseList
	res := s.Client.Query("GET", s.ClientTargetInfo.APIEndpoint, appRESTPath, nil)
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyBytes, &responseList)
	s.AppStats.TotalAppCount = responseList.TotalResults
	return s.AppStats.TotalAppCount
}

//CompileAllApps - compile the information for all of your applications
func (s *AppSearch) CompileAllApps() {
	var responseList cf.APIResponseList
	res := s.Client.Query("GET", s.ClientTargetInfo.APIEndpoint, appRESTPath, url.QueryEscape(fmt.Sprintf("results-per-page=%d", s.AppStats.TotalAppCount)))
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyBytes, &responseList)

	for _, applicationRecord := range responseList.Resources {
		s.processApplicationRecord(applicationRecord)
	}
}

func (s *AppSearch) processApplicationRecord(appRecord cf.APIResponse) {
	s.processAI(appRecord)
	s.processBP(appRecord)
	s.processRunningApps(appRecord)
}

func (s *AppSearch) processAI(appRecord cf.APIResponse) {

	if appRecord.Entity[stateFieldname].(string) == applicationRunningValue {
		s.AppStats.TotalInstanceCount += int(appRecord.Entity[instanceFieldname].(float64))
	}
}

func (s *AppSearch) processRunningApps(appRecord cf.APIResponse) {

	if appRecord.Entity[stateFieldname].(string) == applicationRunningValue {
		s.AppStats.TotalRunningCount++
	}
}

func (s *AppSearch) processBP(appRecord cf.APIResponse) {
	var buildpackFields []string

	if appRecord.Entity[buildpackFieldname] != nil {
		buildpackFields = append(buildpackFields, strings.ToUpper(appRecord.Entity[buildpackFieldname].(string)))
	}

	if appRecord.Entity[detectedBuildpackFieldname] != nil {
		buildpackFields = append(buildpackFields, strings.ToUpper(appRecord.Entity[detectedBuildpackFieldname].(string)))
	}
	s.checkAgainstRegisteredBuildpacks(strings.Join(buildpackFields, " "))
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
