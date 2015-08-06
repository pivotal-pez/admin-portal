package events_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/admin-portal/events"

	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
)

var _ = Describe("EventSearch", func() {
	var (
		fakeClient  *fakeCloudFoundryClient
		eventSearch *EventSearch
	)

	BeforeEach(func() {
		fakeClient = new(fakeCloudFoundryClient)
	})

	Context("Calling CompileRecentEvents successfully", func() {
		BeforeEach(func() {
			eventSearch = new(EventSearch).Init(fakeClient)
			eventSearch.CompileRecentEvents()
		})

		It("should populate eventsblob with a list of event objects", func() {
			controlLength := 17
			controlObject := cf.APIResponse{
				Metadata: cf.APIMetadata{
					GUID:      "73be434e-e703-4ffa-8e8c-40799eaf21a0",
					URL:       "/v2/events/73be434e-e703-4ffa-8e8c-40799eaf21a0",
					CreatedAt: "2015-08-06T16:05:40Z",
					UpdatedAt: "",
				},
				Entity: map[string]interface{}{
					"actee_type":        "app",
					"actee_name":        "adminportal",
					"timestamp":         "2015-08-06T16:05:40Z",
					"space_guid":        "255f7cf9-c855-47fb-af10-e03a0637912f",
					"organization_guid": "45f92610-f258-4f6f-9846-67b1cd98aed3",
					"metadata": map[string]interface{}{
						"request": map[string]interface{}{"state": "STARTED"},
					},
					"type":       "audit.app.update",
					"actor":      "5b3f2caf-68c0-43c4-ab01-b8b1aabb5181",
					"actor_type": "user",
					"actor_name": "system-automate",
					"actee":      "23a775f1-b8c4-4763-9810-560dca003354",
				},
			}
			Ω(len(eventSearch.EventsBlob)).Should(Equal(controlLength))
			Ω(eventSearch.EventsBlob[0]).Should(Equal(controlObject))
		})
	})
})
