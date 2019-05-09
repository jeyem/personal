package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

var html string

const (
	htmlpath   = "views/index.html"
	staticpath = "views/static/"
	content    = "content.json"
	cvpath     = "cv.pdf"
)

func site(c echo.Context) error {
	return c.HTML(200, html)
}

func main() {
	html = parsehtml()
	e := echo.New()
	e.Static("static", staticpath)
	e.GET("/", site)
	e.GET("/cv", func(c echo.Context) error {
		return c.Attachment(cvpath, cvpath)
	})
	e.Logger.Fatal(e.Start(":6060"))
}

func parsehtml() string {
	tpl := pongo2.Must(pongo2.FromFile(htmlpath))
	data, err := ioutil.ReadFile(content)
	if err != nil {
		panic(err)
	}
	params := pongo2.Context{}
	if err := json.Unmarshal(data, &params); err != nil {
		panic(err)
	}
	out, err := tpl.Execute(params)
	if err != nil {
		panic(err)
	}
	return out
}
