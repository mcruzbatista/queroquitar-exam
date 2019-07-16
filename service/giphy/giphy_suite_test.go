package giphy_test

import (
	"github.com/jarcoal/httpmock"
	"gopkg.in/resty.v1"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGiphy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Giphy Suite")
}

var _ = BeforeSuite(func() {
	// block all HTTP requests
	httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
})

var _ = BeforeEach(func() {
	// remove any mocks
	httpmock.Reset()
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})