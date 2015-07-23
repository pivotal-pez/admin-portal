package cloudfoundryclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

var _ = Describe("CFClient", func() {
	Describe("Query", func() {
		var (
			cfclient   CloudFoundryClient
			controlRes = mockHttpResponse(mockSuccessUserResponseBody, 200)
		)

		Context("Query called", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: controlRes,
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should just pass back the response object", func() {
				res := cfclient.Query("GET", "mydomain.com", "/v2/User", "")
				Ω(res).Should(Equal(controlRes))
			})
		})
	})

	Describe("QueryUsers", func() {
		var (
			cfclient         CloudFoundryClient
			controlResources = UserAPIResponse{
				Schemas: []string{
					"urn:scim:schemas:core:1.0",
				},
				StartIndex:   1,
				ItemsPerPage: 100,
				TotalResults: 1,
				Resources: []UserResource{
					UserResource{
						Active:    false,
						Approvals: nil,
						Emails:    nil,
						Groups:    nil,
						ID:        "123456",
						Meta:      nil,
						Name:      nil,
						Origin:    "",
						Schemas:   nil,
						UserName:  "testuser",
						Verified:  false,
					},
				},
			}
		)

		Context("QueryUsers called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessUserResponseBody, mockSuccessUserStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				users, err := cfclient.QueryUsers(1, 1, "id,userName", "")
				Ω(err).Should(BeNil())
				Ω(users).Should(BeEquivalentTo(controlResources))
			})

			It("should parse the response object without error without any attributes", func() {
				users, err := cfclient.QueryUsers(1, 1, "", "")
				Ω(err).Should(BeNil())
				Ω(users).Should(BeEquivalentTo(controlResources))
			})
		})

		Context("QueryUsers unsuccessful response", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessUserResponseBody, (mockSuccessUserStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should return an error", func() {
				users, err := cfclient.QueryUsers(1, 1, "id,userName", "")
				Ω(err).Should(Equal(ErrFailedStatusCode))
				Ω(users).Should(Equal(UserAPIResponse{}))
			})
		})
	})

	Describe("AddSpace", func() {
		var (
			cfclient         CloudFoundryClient
			controlOrgGUID   = "ca8a7fb0-737a-4fe9-8b28-42b064981abe"
			controlSpaceGUID = "da840d5b-7987-4c6e-8c3c-b0ebead3b4ed"
			controlSpaceName = "development"
		)

		Context("AddSpace called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessSpaceResponseBody, mockSuccessSpaceStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				guid, err := cfclient.AddSpace(controlSpaceName, controlOrgGUID)
				Ω(err).Should(BeNil())
				Ω(guid).Should(Equal(controlSpaceGUID))
			})
		})

		Context("AddSpace unsuccessful response", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessSpaceResponseBody, (mockSuccessSpaceStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should return an error", func() {
				guid, err := cfclient.AddSpace(controlSpaceName, controlOrgGUID)
				Ω(err).Should(Equal(ErrSpaceCreateAPICallFailure))
				Ω(guid).Should(BeEmpty())
			})
		})
	})

	Describe("RemoveOrg", func() {
		var (
			cfclient       CloudFoundryClient
			controlOrgGUID = "1e2bae2c-459e-4ad8-b1cb-ffc09d209b32"
		)

		Context("RemoveOrg called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse("", mockSuccessRemoveOrgStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				err := cfclient.RemoveOrg(controlOrgGUID)
				Ω(err).Should(BeNil())
			})
		})

		Context("RemoveOrg unsuccessful response", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse("", (mockSuccessRemoveOrgStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should return an error", func() {
				err := cfclient.RemoveOrg(controlOrgGUID)
				Ω(err).Should(Equal(ErrOrgRemoveAPICallFailure))
			})
		})
	})

	Describe("AddOrg", func() {
		var (
			cfclient       CloudFoundryClient
			controlOrgGUID = "1e2bae2c-459e-4ad8-b1cb-ffc09d209b32"
			controlOrgName = "my-org-name"
		)

		Context("AddOrg called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessOrgResponseBody, mockSuccessOrgStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				guid, err := cfclient.AddOrg(controlOrgGUID)
				Ω(err).Should(BeNil())
				Ω(guid).Should(Equal(controlOrgGUID))
			})
		})

		Context("AddOrg unsuccessful response", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessOrgResponseBody, (mockSuccessOrgStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should return an error", func() {
				guid, err := cfclient.AddOrg(controlOrgName)
				Ω(err).Should(Equal(ErrOrgCreateAPICallFailure))
				Ω(guid).Should(BeEmpty())
			})
		})
	})

	Describe("AddRole", func() {
		var cfclient CloudFoundryClient

		Context("AddRole called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessRoleResponseBody, mockSuccessRoleStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				err := cfclient.AddRole(OrgEndpoint, "target-guid-12345", RoleTypeManager, "user-guid-12345")
				Ω(err).Should(BeNil())
			})
		})

		Context("AddRole unsuccessful response", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessRoleResponseBody, (mockSuccessRoleStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should return an error", func() {
				err := cfclient.AddRole(OrgEndpoint, "target-guid-12345", RoleTypeManager, "user-guid-12345")
				Ω(err).Should(Equal(ErrFailedStatusCode))
			})
		})
	})

	Describe("QueryUserGUID", func() {
		var cfclient CloudFoundryClient

		Context("QueryUserGUID called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessUserResponseBody, mockSuccessUserStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				controlUser := "testuser"
				controlUID := "123456"
				guid, err := cfclient.QueryUserGUID(controlUser)
				Ω(guid).Should(Equal(controlUID))
				Ω(err).Should(BeNil())
			})
		})

		Context("QueryUserGUID called w/ invalid user", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse("{}", mockSuccessUserStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the return a user not found error", func() {
				controlUser := "invalid-user"
				guid, err := cfclient.QueryUserGUID(controlUser)
				Ω(guid).Should(BeEmpty())
				Ω(err).Should(Equal(ErrNoUserFound))
			})
		})

		Context("QueryUserGUID call failed", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessUserResponseBody, (mockSuccessUserStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the return a user not found error", func() {
				controlUser := "testuser"
				guid, err := cfclient.QueryUserGUID(controlUser)
				Ω(guid).Should(BeEmpty())
				Ω(err).Should(Equal(ErrFailedStatusCode))
			})
		})
	})

	Describe("QueryAPIInfo", func() {
		var cfclient CloudFoundryClient

		Context("QueryAPIInfo called successfully", func() {
			var controlAPIDomain = "api.test.org"

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessInfoResponseBody, mockSuccessInfoStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer:        mockDoer,
					apiEndpoint: controlAPIDomain,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				info, err := cfclient.QueryAPIInfo()
				Ω(info.LoggingEndpoint).ShouldNot(BeEmpty())
				Ω(info.AuthorizationEndpoint).ShouldNot(BeEmpty())
				Ω(info.TokenEndpoint).ShouldNot(BeEmpty())
				Ω(info.APIEndpoint).Should(Equal(controlAPIDomain))
				Ω(err).Should(BeNil())
			})
		})

		Context("QueryAPIInfo called with failure", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessInfoResponseBody, (mockSuccessInfoStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				info, err := cfclient.QueryAPIInfo()
				Ω(info.LoggingEndpoint).Should(BeEmpty())
				Ω(info.AuthorizationEndpoint).Should(BeEmpty())
				Ω(info.TokenEndpoint).Should(BeEmpty())
				Ω(err).ShouldNot(BeNil())
			})
		})
	})
})
