package main

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/sc0rp1us/dotviz-server/controllers"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("public"))
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
	}))

	dotController := new(controllers.DotController)
	m.Get("/dot/:format/:data", dotController.Render)

	m.Run()
}
