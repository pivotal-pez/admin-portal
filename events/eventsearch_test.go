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
			controlLength := 50
			controlObject := cf.APIResponse{
				Metadata: cf.APIMetadata{
					GUID:      "782c711f-ed87-487d-ab53-9d4f17d53b6a",
					URL:       "/v2/events/782c711f-ed87-487d-ab53-9d4f17d53b6a",
					CreatedAt: "2015-07-06T20:00:39Z",
					UpdatedAt: "",
				},
				Entity: map[string]interface{}{
					"actor":      "1a3cf3ba-dad6-4c92-981f-b58e0f939aa5",
					"actor_name": "jcalabrese@pivotal.io",
					"metadata": map[string]interface{}{
						"request": map[string]interface{}{"name": "hcfdev"},
					},
					"actee_name":        "hcfdev",
					"actee_type":        "app",
					"actor_type":        "user",
					"organization_guid": "45f92610-f258-4f6f-9846-67b1cd98aed3",
					"space_guid":        "ea88ed9e-91f1-4763-8eef-54fe38acf603",
					"timestamp":         "2015-07-06T20:00:39Z",
					"type":              "audit.app.update",
					"actee":             "a8e04af5-4e9a-465a-a924-08561cff964d",
				},
			}
			Ω(len(eventSearch.EventsBlob)).Should(Equal(controlLength))
			Ω(eventSearch.EventsBlob[0]).Should(Equal(controlObject))
		})
	})
})
