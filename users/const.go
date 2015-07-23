package users

import "errors"

var (
	SpacesSuccessStatusCode = 200
	ErrListUserSpaces       = errors.New("there was an error calling list user spaces endpoint")
	OrgsSuccessStatusCode   = 200
	ErrListUserOrgs         = errors.New("there was an error calling list user orgs endpoint")
)
