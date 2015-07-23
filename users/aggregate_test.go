package users_test

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"

	. "github.com/pivotal-pez/admin-portal/users"
)

var _ = Describe("UserAggregate", func() {

	Context("Calling .Compile() with a user response list from uaa server", func() {
		controlUserBlob := &UserAggregate{
			UserCount:     321,
			UAACount:      72,
			ExternalCount: 249,
			OrphanedCount: 0,
		}

		It("should aggregate those results", func() {
			userBlob := new(UserAggregate)
			contents, _ := ioutil.ReadFile("fixtures/userlist.json")
			response := cf.UserAPIResponse{}
			json.Unmarshal(contents, &response)
			userBlob.Compile(response)
			Ω(userBlob).Should(Equal(controlUserBlob))
		})
	})
})
