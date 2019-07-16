package html

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"queroquitar-exam/service/giphy"
)

type Html interface {
	GetHtmlGiphy(c echo.Context) error
}

type htmlResource struct {}

func (hr *htmlResource) GetHtmlGiphy(c echo.Context) (err error) {

	gifUrl,err := giphy.GetRandomGiphy()

	if err != nil {
		logrus.Error(err.Error())
		return c.Render(http.StatusBadRequest,"error","")
	}

	return c.Render(http.StatusOK, "giphy", gifUrl)
}

func NewHtml() Html {
	return &htmlResource{}
}