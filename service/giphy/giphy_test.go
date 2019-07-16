package giphy_test

import (
	"errors"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"queroquitar-exam/service/giphy"
)

var _ = Describe("Giphy", func() {
	Context("Success GetRandomGiphy", func() {
		It("call GetRandomGiphy and return random giphy and nil error", func() {
			//Mock http
			body :=`{"data":{"image_url":"https://any.gif.image/queroquitar.gif"}`
			responder := httpmock.NewStringResponder(200, body)
			url := "https://api.giphy.com/v1/gifs/random?api_key=dGiRPc2btcRacSHPVKR4V7NzR47Shliw&rating=G"
			httpmock.RegisterResponder("GET",url, responder)

			//call func
			gifUrl, err := giphy.GetRandomGiphy()

			Expect(gifUrl).ToNot(BeNil())
			Expect(gifUrl).To(Equal("https://any.gif.image/queroquitar.gif"))
			Expect(err).To(BeNil())
		})
	})

	Context("Error GetRandomGiphy", func() {
		It("call GetRandomGiphy and return error from resty request", func() {
			//Mock http
			responder := httpmock.NewErrorResponder(errors.New("error on something"))
			url := "https://api.giphy.com/v1/gifs/random?api_key=dGiRPc2btcRacSHPVKR4V7NzR47Shliw&rating=G"

			httpmock.RegisterResponder("GET",url, responder)

			//call func
			gifUrl, err := giphy.GetRandomGiphy()

			Expect(gifUrl).To(Equal(""))
			Expect(err.Error()).To(Equal("Error on Resty:Get https://api.giphy.com/v1/gifs/random?api_key=dGiRPc2btcRacSHPVKR4V7NzR47Shliw&rating=G: error on something"))
		})
	})

	Context("Error GetRandomGiphy", func() {
		It("call GetRandomGiphy and return error from empty gifUrl", func() {
			//Mock http
			body :=`{"data":{"image_url":""}`
			responder := httpmock.NewStringResponder(200, body)
			url := "https://api.giphy.com/v1/gifs/random?api_key=dGiRPc2btcRacSHPVKR4V7NzR47Shliw&rating=G"

			httpmock.RegisterResponder("GET",url, responder)

			//call func
			gifUrl, err := giphy.GetRandomGiphy()

			Expect(gifUrl).To(Equal(""))
			Expect(err.Error()).To(Equal("no giphy found"))
		})
	})
})
