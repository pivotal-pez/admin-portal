package applications_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/admin-portal/applications"
)

var _ = Describe("AppSearch", func() {

	Context("Calling ListAppCount successfully", func() {
		appSearch := new(AppSearch).Init(new(fakeCloudFoundryClient))

		It("should yield the number of apps that exist in the foundation", func() {
			controlAppCount := 13
			Ω(appSearch.GetAppCount()).Should(Equal(controlAppCount))
			Ω(appSearch.AppStats.TotalAppCount).Should(Equal(controlAppCount))
		})
	})

	Context("Calling CompileAllApps successfully", func() {
		appSearch := new(AppSearch).Init(new(fakeCloudFoundryClient))
		appSearch.GetAppCount()
		appSearch.CompileAllApps()

		It("should log the proper instance count", func() {
			controlAICount := 3
			Ω(appSearch.AppStats.TotalInstanceCount).Should(Equal(controlAICount))
		})

		It("should log the proper running application count", func() {
			controlRunningCount := 2
			Ω(appSearch.AppStats.TotalRunningCount).Should(Equal(controlRunningCount))
		})

		It("should log the proper running memory usage", func() {
			controlMemory := float64(3)
			Ω(appSearch.AppStats.TotalMemory).Should(Equal(controlMemory))
		})

		It("should log the proper Buildpack count", func() {
			controlJBPCount := 1
			controlRubyCount := 2
			controlGOCount := 1
			controlPythonCount := 1
			controlPHPCount := 1
			controlNodeCount := 1
			controlOtherCount := 1
			Ω(appSearch.AppStats.JavaBPCount).Should(Equal(controlJBPCount))
			Ω(appSearch.AppStats.RubyBPCount).Should(Equal(controlRubyCount))
			Ω(appSearch.AppStats.NodeBPCount).Should(Equal(controlNodeCount))
			Ω(appSearch.AppStats.GOBPCount).Should(Equal(controlGOCount))
			Ω(appSearch.AppStats.PythonBPCount).Should(Equal(controlPythonCount))
			Ω(appSearch.AppStats.PHPBPCount).Should(Equal(controlPHPCount))
			Ω(appSearch.AppStats.OtherBPCount).Should(Equal(controlOtherCount))
		})
	})
})
