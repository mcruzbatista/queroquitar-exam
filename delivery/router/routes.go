package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"queroquitar-exam/delivery/html"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Routes() {

	h := html.NewHtml()

	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.gohtml")),
	}

	e := echo.New()

	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/giphy", h.GetHtmlGiphy)

	e.Logger.Fatal(e.Start(":8080"))
}
