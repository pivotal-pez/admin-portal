package users_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"

	"github.com/pivotal-pez/admin-portal/users"
)

func compileFromFile(filename string) (userBlob *users.UserAggregate) {
	userBlob = new(users.UserAggregate)
	contents, _ := ioutil.ReadFile("fixtures/userlist.json")
	response := cf.UserAPIResponse{}
	json.Unmarshal(contents, &response)
	userBlob.Compile(response)
	return
}

func sortMapKeys(m map[string]int) (keys []string) {

	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return
}

var _ = Describe("UserAggregate", func() {
	controlEmptyCreateBucket := map[string]int{
		"2015-07-05": 0,
		"2015-07-16": 4,
		"2015-07-17": 1,
		"2015-07-20": 11,
		"2015-07-08": 4,
		"2015-07-11": 0,
		"2015-07-19": 0,
		"2015-07-01": 2,
		"2015-07-04": 0,
		"2015-07-06": 2,
		"2015-07-07": 3,
		"2015-07-10": 3,
		"2015-07-13": 0,
		"2015-07-15": 6,
		"2015-07-03": 1,
		"2015-07-02": 2,
		"2015-07-09": 4,
		"2015-07-12": 1,
		"2015-07-14": 5,
		"2015-07-18": 1,
		"2015-06-30": 0,
	}

	BeforeEach(func() {
		users.GetCurrentDate = func() time.Time {
			loc, _ := time.LoadLocation("")
			return time.Date(2015, time.July, 20, 0, 0, 0, 0, loc)
		}
	})

	Describe("TimeStringifier", func() {
		Context("when passed a time.Time", func() {
			It("should return a string of just the YYYY-MM-DD", func() {
				year := 2012
				day := 12
				month := time.April
				controlTimeString := fmt.Sprintf("%d-%02d-%02d", year, month, day)
				location, _ := time.LoadLocation("")
				sampleDate := time.Date(year, month, day, 0, 0, 0, 0, location)
				Ω(controlTimeString).Should(Equal(users.TimeStringifier(sampleDate)))
			})
		})
	})

	Describe(".GenerateUserCreateBuckets()", func() {
		Context("when called", func() {

			controlUserBlob := &users.UserAggregate{
				CreateDayOverDay: controlEmptyCreateBucket,
			}

			It("should initialize the date buckets properly", func() {
				userBlob := compileFromFile("fixtures/userlist.json")
				Ω(sortMapKeys(controlUserBlob.CreateDayOverDay)).Should(Equal(sortMapKeys(userBlob.CreateDayOverDay)))
			})
		})
	})

	Context("Calling .Compile() with a user response list from uaa server", func() {

		controlUserBlob := &users.UserAggregate{
			UserCount:        321,
			UAACount:         72,
			ExternalCount:    249,
			OrphanedCount:    0,
			CreateDayOverDay: controlEmptyCreateBucket,
		}

		It("should aggregate those results", func() {
			userBlob := compileFromFile("fixtures/userlist.json")
			Ω(userBlob.UserCount).Should(Equal(controlUserBlob.UserCount))
			Ω(userBlob.UAACount).Should(Equal(controlUserBlob.UAACount))
			Ω(userBlob.ExternalCount).Should(Equal(controlUserBlob.ExternalCount))
			Ω(userBlob.OrphanedCount).Should(Equal(controlUserBlob.OrphanedCount))
			Ω(userBlob.CreateDayOverDay).Should(Equal(controlUserBlob.CreateDayOverDay))
		})
	})
})
