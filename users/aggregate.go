package users

import (
	"fmt"
	"time"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

//UserAggregate - a user info aggregation object. bucket of stuff
type (
	UserAggregate struct {
		UserCount        int
		UAACount         int
		ExternalCount    int
		OrphanedCount    int
		CreateDayOverDay map[string]int
	}
)

var GetCurrentDate = func() time.Time {
	return time.Now()
}

//Compile - compile the info aggregates
func (s *UserAggregate) Compile(users cf.UserAPIResponse) {
	s.GenerateUserCreateBuckets()
	s.UserCount = users.TotalResults

	for _, v := range users.Resources {

		if v.Origin == InvitedGuestUserValue {
			s.UAACount++

		} else {
			s.ExternalCount++
		}
		s.compileCreateUserAggregate(v.Meta[CreatedFieldname].(string)[:10])
	}
}

func (s *UserAggregate) compileCreateUserAggregate(createDate string) {

	if _, ok := s.CreateDayOverDay[createDate]; ok {
		s.CreateDayOverDay[createDate]++
	}
}

func (s *UserAggregate) GenerateUserCreateBuckets() {
	now := GetCurrentDate()
	s.CreateDayOverDay = make(map[string]int)
	hash := TimeStringifier(now)
	s.CreateDayOverDay[hash] = 0

	for i := DayOverDayHistoryLimit; i >= 0; i-- {
		t := now.AddDate(0, 0, (i * -1))
		hash := TimeStringifier(t)
		s.CreateDayOverDay[hash] = 0
	}
}

func TimeStringifier(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}
