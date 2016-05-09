package cloudfoundryclient

import (
	"errors"
	"net/http"
)

const (
	//OrgCreateSuccessStatusCode - success status code from a call to the org create cc endpoint
	OrgCreateSuccessStatusCode = http.StatusCreated
	//OrgEndpoint - the endpoint to hit for org actions
	OrgEndpoint = "/v2/organizations"
	//SpacesEndpont - the endpoint to hit for spaces actions
	SpacesEndpont = "/v2/spaces"
	//SpacesCreateSuccessStatusCode = success status code of spaces rest call
	SpacesCreateSuccessStatusCode = http.StatusCreated
	//ListUsersEndpoint - get a list of all users in paas
	ListUsersEndpoint = "/Users"
	//ListUsersSuccessStatus - success status code for users call
	ListUsersSuccessStatus = http.StatusOK
	//InfoURLPath - the endpoint to grab api info data
	InfoURLPath = "/v2/info"
	//InfoSuccessStatus --
	InfoSuccessStatus = http.StatusOK
	//RoleTypeManager - this is the managers type for role assignments
	RoleTypeManager = "managers"
	//RoleTypeUser - this is the users type for role assignments
	RoleTypeUser = "users"
	//RoleTypeDeveloper - a role type for developers of a space
	RoleTypeDeveloper = "developers"
	//RoleCreationURLFormat - formatter string for role creation url generation
	RoleCreationURLFormat = "%s/%s/%s/%s"
	//RoleCreateSuccessStatusCode - success status code for role assignment calls
	RoleCreateSuccessStatusCode = http.StatusCreated
	//OrgRemoveSuccessStatus - success status code for org removal
	OrgRemoveSuccessStatus = http.StatusNoContent
)

var (
	//ErrOrgCreateAPICallFailure - error for failed call to create org endpoint
	ErrOrgCreateAPICallFailure = errors.New("failed to create org on api call")
	//ErrOrgRemoveAPICallFailure - error for failed call to remove org endpoint
	ErrOrgRemoveAPICallFailure = errors.New("failed to remove org on api call")
	//ErrSpaceCreateAPICallFailure - error for failed call to create org endpoint
	ErrSpaceCreateAPICallFailure = errors.New("failed to create space on api call")
	//ErrNoUserFound - error no user found
	ErrNoUserFound = errors.New("no matching user found in system")
	//ErrFailedStatusCode - we recieved a status code not matching the success code for the endpoint
	ErrFailedStatusCode = errors.New("status code response does not match the known success status code for rest endpoint")
)
