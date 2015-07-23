package users

import (
	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

//UserAggregate - a user info aggregation object. bucket of stuff
type UserAggregate struct {
	UserCount     int
	UAACount      int
	ExternalCount int
	OrphanedCount int
}

//Compile - compile the info aggregates
func (s *UserAggregate) Compile(users cf.UserAPIResponse) {
	s.UserCount = users.TotalResults

	for _, v := range users.Resources {

		if v.Origin == "uaa" {
			s.UAACount++

		} else {
			s.ExternalCount++
		}
	}
}
