package main

import (
	"html/template"
	"io"

	"github.com/aruga-dev/arugaONE-API/util/consts"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	h := &AppHandler{}

	e := echo.New()

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.Use(middleware.Recover())
	e.POST("/push", h.PushMessage)
	e.POST("/webhook/:botName", h.Webhook)
	e.GET("/liff", h.LiffPage)
	e.POST("/liff", h.LiffSubmit)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		e.DefaultHTTPErrorHandler(err, c)
		print(err.Error())
	}

	if !consts.IsLocal() {
		e.Pre(middleware.HTTPSWWWRedirect())
		e.StartTLS(":443", consts.TLSCertificatePath(), consts.TLSPrivateKeyPath())
	} else {
		e.Start(":80")
	}
}
