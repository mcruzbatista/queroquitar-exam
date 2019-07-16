package html_test

import (
	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"queroquitar-exam/delivery/html"
)

var _ = Describe("Html", func() {
	Context("Success GetRandomGiphy", func() {
		It("call GetRandomGiphy and return code 200", func() {
			//Mock http
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/giphy")

			h := html.NewHtml()
			h.GetHtmlGiphy(c)
			Expect(rec.Code).To(Equal(http.StatusOK))
		})
	})
})

