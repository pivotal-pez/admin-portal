package cloudfoundryclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//NewCloudFoundryClient - generate a new cloudfoundryclient interface object
func NewCloudFoundryClient(auth AuthRequestCreator, log logger) CloudFoundryClient {
	c := &CFClient{
		RequestDecorator: auth,
		Log:              log,
		Info:             new(CloudFoundryAPIInfo),
	}
	return c
}

//QueryAPIInfo - get the info results for your target rest api
func (s *CFClient) QueryAPIInfo() (info *CloudFoundryAPIInfo, err error) {
	rest := &RestRunner{
		Logger:            s.Log,
		RequestDecorator:  s.RequestDecorator,
		Verb:              "GET",
		URL:               s.RequestDecorator.CCTarget(),
		Path:              InfoURLPath,
		SuccessStatusCode: InfoSuccessStatus,
	}
	rest.OnSuccess = func(res *http.Response) {
		b, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(b, &s.Info)
		s.Info.APIEndpoint = s.RequestDecorator.CCTarget()
		s.Log.Println(s.Info)
		s.Log.Println(string(b[:]))
	}
	rest.OnFailure = func(res *http.Response, e error) {
		b, _ := ioutil.ReadAll(res.Body)
		s.Log.Println("error: ", e, string(b[:]))
		err = e
	}
	rest.Run()
	info = s.Info
	return
}

//Query - make a generic query against any rest endpoint
func (s *CFClient) Query(verb string, domain string, path string, args interface{}) (response *http.Response) {
	rest := &RestRunner{
		Logger:            s.Log,
		RequestDecorator:  s.RequestDecorator,
		Verb:              verb,
		URL:               domain,
		Path:              path,
		SuccessStatusCode: 0,
		Data:              args,
	}
	rest.OnSuccess = func(res *http.Response) {
		response = res
	}
	rest.OnFailure = func(res *http.Response, e error) {
		response = res
	}
	rest.Run()
	return
}

//QueryUsers - get the guid for the given user
func (s *CFClient) QueryUsers(startIndex, count int, attributes, filter string) (userList UserAPIResponse, err error) {
	var (
		data = fmt.Sprintf("startIndex=%d&count=%d%s%s", startIndex, count, parseArg("attributes", attributes), parseArg("filter", filter))
	)

	rest := &RestRunner{
		Logger:            s.Log,
		RequestDecorator:  s.RequestDecorator,
		Verb:              "GET",
		URL:               s.Info.TokenEndpoint,
		Path:              ListUsersEndpoint,
		SuccessStatusCode: ListUsersSuccessStatus,
		Data:              data,
	}
	rest.OnSuccess = func(res *http.Response) {
		s.Log.Println("user response: ", res)
		b, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(b, &userList)
	}
	rest.OnFailure = func(res *http.Response, e error) {
		b, _ := ioutil.ReadAll(res.Body)
		s.Log.Println("call for user guid failed :(", e, string(b[:]))
		err = e
	}
	rest.Run()
	return
}

func parseArg(name string, filter string) (parsedFilter string) {

	if filter == "" {
		parsedFilter = filter
	} else {
		parsedFilter = fmt.Sprintf("&%s=%s", name, filter)
	}
	return
}

//QueryUserGUID - get the guid for the given user
func (s *CFClient) QueryUserGUID(username string) (guid string, err error) {
	var (
		users      = UserAPIResponse{}
		userFilter = url.QueryEscape(fmt.Sprintf("userName eq '%s'", username))
	)

	if users, err = s.QueryUsers(1, 1, "id", userFilter); err == nil && users.TotalResults > 0 {
		guid = users.Resources[0].ID

	} else {

		if err == nil {
			err = ErrNoUserFound
		}
	}
	return
}

//AddRole - add a role mapping a user to a org or space
func (s *CFClient) AddRole(rolePathPrefix string, targetGUID string, roleType string, userGUID string) (err error) {
	rolePath := fmt.Sprintf(RoleCreationURLFormat, rolePathPrefix, targetGUID, roleType, userGUID)
	rest := &RestRunner{
		Logger:            s.Log,
		RequestDecorator:  s.RequestDecorator,
		Verb:              "PUT",
		URL:               s.RequestDecorator.CCTarget(),
		Path:              rolePath,
		SuccessStatusCode: RoleCreateSuccessStatusCode,
	}
	rest.OnSuccess = func(res *http.Response) {
		s.Log.Println("we have a role!", rolePath)
	}
	rest.OnFailure = func(res *http.Response, e error) {
		b, _ := ioutil.ReadAll(res.Body)
		s.Log.Println("call for role failed :(", rolePath, e, res.StatusCode, string(b[:]))
		err = e
	}
	rest.Run()
	return
}

//AddOrg - add an org to the target foundation
func (s *CFClient) AddOrg(orgName string) (orgGUID string, err error) {
	var (
		data = fmt.Sprintf(`{"name":"%s"}`, orgName)
	)
	rest := &RestRunner{
		Logger:            s.Log,
		RequestDecorator:  s.RequestDecorator,
		Verb:              "POST",
		URL:               s.RequestDecorator.CCTarget(),
		Path:              OrgEndpoint,
		SuccessStatusCode: OrgCreateSuccessStatusCode,
		Data:              data,
	}
	rest.OnSuccess = func(res *http.Response) {
		s.Log.Println("we created the org successfully")
		apiResponse := new(APIResponse)
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, apiResponse)
		orgGUID = apiResponse.Metadata.GUID
	}
	rest.OnFailure = func(res *http.Response, e error) {
		body, _ := ioutil.ReadAll(res.Body)
		logMessage := fmt.Sprintf("call to create org api failed: error-(%s) responsebody-(%s)", e.Error(), string(body[:]))
		s.Log.Println(logMessage)
		err = ErrOrgCreateAPICallFailure
	}
	rest.Run()
	return
}

//RemoveOrg - remove a org by guid
func (s *CFClient) RemoveOrg(orgGUID string) (err error) {
	orgDeletePath := fmt.Sprintf("%s/%s", OrgEndpoint, orgGUID)
	rest := &RestRunner{
		Logger:            s.Log,
		RequestDecorator:  s.RequestDecorator,
		Verb:              "DELETE",
		URL:               s.RequestDecorator.CCTarget(),
		Path:              orgDeletePath,
		SuccessStatusCode: OrgRemoveSuccessStatus,
		Data:              "",
	}
	rest.OnSuccess = func(res *http.Response) {
		s.Log.Println("we removed the org successfully")
	}
	rest.OnFailure = func(res *http.Response, e error) {
		s.Log.Println("call to remove org api failed")
		err = ErrOrgRemoveAPICallFailure
	}
	rest.Run()
	return

	return
}

//AddSpace - add a space to the given org
func (s *CFClient) AddSpace(spaceName string, orgGUID string) (spaceGUID string, err error) {
	var (
		data = fmt.Sprintf(`{"name": "%s","organization_guid":"%s"}`, spaceName, orgGUID)
	)
	rest := &RestRunner{
		Logger:            s.Log,
		RequestDecorator:  s.RequestDecorator,
		Verb:              "POST",
		URL:               s.RequestDecorator.CCTarget(),
		Path:              SpacesEndpont,
		SuccessStatusCode: SpacesCreateSuccessStatusCode,
		Data:              data,
	}
	rest.OnSuccess = func(res *http.Response) {
		s.Log.Println("we created the space successfully")
		apiResponse := new(APIResponse)
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, apiResponse)
		spaceGUID = apiResponse.Metadata.GUID
	}
	rest.OnFailure = func(res *http.Response, e error) {
		s.Log.Println("call to create space api failed")
		err = ErrSpaceCreateAPICallFailure
	}
	rest.Run()
	return
}

//AddUser - add a user to the targetted foundation
func (s *CFClient) AddUser(username string) (err error) {
	return
}
