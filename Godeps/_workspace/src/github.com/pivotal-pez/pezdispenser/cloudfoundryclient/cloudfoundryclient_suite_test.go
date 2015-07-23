package cloudfoundryclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCloudFoundryClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cloud Foundry Client Suite")
}
