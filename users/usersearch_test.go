package users_test

import (
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/pivotal-pez/admin-portal/users"
	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

type fakeCloudFoundryClient struct {
	QueryUsersErr error
	QueryUsersRes cf.UserAPIResponse
}

func (s *fakeCloudFoundryClient) QueryAPIInfo() (info *cf.CloudFoundryAPIInfo, err error) {
	return
}

func (s *fakeCloudFoundryClient) QueryUsers(start int, end int, filter string, search string) (userList cf.UserAPIResponse, err error) {
	return s.QueryUsersRes, s.QueryUsersErr
}

func (s *fakeCloudFoundryClient) Query(verb string, domain string, path string, args interface{}) (response *http.Response) {
	return
}

var _ = Describe("UserSearch", func() {
	var (
		userSearch *UserSearch
	)

	BeforeEach(func() {
		userSearch = new(UserSearch)
		userSearch.Client = new(fakeCloudFoundryClient)
	})

	Describe(".BuildQuery()", func() {
		Context("Calling with non blank arguments", func() {
			It("should return a properly structured query string", func() {
				controlUserType := "usertype"
				controlUserName := "username"
				controlResponse := "origin+eq+%27usertype%27+and+userName+co+%27username%27"
				Ω(userSearch.BuildQuery(controlUserType, controlUserName)).Should(Equal(controlResponse))
			})
		})

		Context("Calling with one blank argument", func() {
			It("should return a properly formatted string", func() {
				controlUserType := ""
				controlUserName := "username"
				controlResponse := "userName+co+%27username%27"
				Ω(userSearch.BuildQuery(controlUserType, controlUserName)).Should(Equal(controlResponse))
			})
		})

		Context("Calling with blank arguments", func() {
			It("should return a properly formatted string", func() {
				controlUserType := ""
				controlUserName := ""
				controlResponse := ""
				Ω(userSearch.BuildQuery(controlUserType, controlUserName)).Should(Equal(controlResponse))
			})
		})
	})

	Describe(".List()", func() {

		Context("calling list with a usertype & username", func() {
			controlUserResponse := cf.UserAPIResponse{}

			BeforeEach(func() {
				userSearch.Client = &fakeCloudFoundryClient{
					QueryUsersErr: nil,
					QueryUsersRes: controlUserResponse,
				}
			})

			It("should return a list of matches", func() {
				controlUserType := ""
				controlUserName := ""
				users, err := userSearch.List(controlUserType, controlUserName)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(users).Should(Equal(controlUserResponse))
			})
		})

		Context("call queries in error", func() {
			controlErr := errors.New("my fake error")

			BeforeEach(func() {
				userSearch.Client = &fakeCloudFoundryClient{
					QueryUsersErr: controlErr,
				}
			})

			It("should return the query error", func() {
				controlUserType := ""
				controlUserName := ""
				_, err := userSearch.List(controlUserType, controlUserName)
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	XDescribe(".ListUserSpaces()", func() {

		Context("calling list with a userGUID", func() {
			controlUserResponse := cf.UserAPIResponse{}

			BeforeEach(func() {
				userSearch.Client = &fakeCloudFoundryClient{
					QueryUsersErr: nil,
					QueryUsersRes: controlUserResponse,
				}
			})

			It("should return a list of matches", func() {
				controlGUID := ""
				users, err := userSearch.ListUserOrgs(controlGUID)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(users).Should(Equal(controlUserResponse))
			})
		})

		Context("call queries in error", func() {
			controlErr := errors.New("my fake error")

			BeforeEach(func() {
				userSearch.Client = &fakeCloudFoundryClient{
					QueryUsersErr: controlErr,
				}
			})

			It("should return the query error", func() {
				controlGUID := ""
				_, err := userSearch.ListUserOrgs(controlGUID)
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	XDescribe(".ListUserOrgs()", func() {

		Context("calling list with a userGUID", func() {
			controlUserResponse := cf.UserAPIResponse{}

			BeforeEach(func() {
				userSearch.Client = &fakeCloudFoundryClient{
					QueryUsersErr: nil,
					QueryUsersRes: controlUserResponse,
				}
			})

			It("should return a list of matches", func() {
				controlGUID := ""
				users, err := userSearch.ListUserOrgs(controlGUID)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(users).Should(Equal(controlUserResponse))
			})
		})

		Context("call queries in error", func() {
			controlErr := errors.New("my fake error")

			BeforeEach(func() {
				userSearch.Client = &fakeCloudFoundryClient{
					QueryUsersErr: controlErr,
				}
			})

			It("should return the query error", func() {
				controlGUID := ""
				_, err := userSearch.ListUserOrgs(controlGUID)
				Ω(err).Should(HaveOccurred())
			})
		})
	})

})
