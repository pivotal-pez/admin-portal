package users

import "errors"

var (
	//SpacesSuccessStatusCode - return success status for a space api call
	SpacesSuccessStatusCode = 200
	//ErrListUserSpaces - error when list users on spaces fails
	ErrListUserSpaces = errors.New("there was an error calling list user spaces endpoint")
	//OrgsSuccessStatusCode - org api call success status code
	OrgsSuccessStatusCode = 200
	//ErrListUserOrgs - error object for failed user org request
	ErrListUserOrgs = errors.New("there was an error calling list user orgs endpoint")
	//InvitedGuestUserValue - uaa record value for a invited user
	InvitedGuestUserValue = "uaa"
	//CreatedFieldname - uaa record fieldname for a user creation timestamp
	CreatedFieldname = "created"
	//DayOverDayHistoryLimit - limit for how far back the day over day user lookup should go
	DayOverDayHistoryLimit = 20
)
