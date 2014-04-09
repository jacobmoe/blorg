package server

import (
	"github.com/go-martini/martini"
	// "github.com/martini-contrib/render"
	// "github.com/jacobmoe/blorg/helper"
	"github.com/jacobmoe/blorg/helper"
	"github.com/jacobmoe/render"
	"html/template"
)

func Run() {
	server := martini.Classic()

	// setup renderer middleware
	server.Use(render.Renderer(render.Options{
		Layout:     "layouts/application",
		IndentJSON: true,
		Funcs:      []template.FuncMap{helper.Application()},
	}))

	// serve static assets from "assets" in addition to "public"
	server.Use(martini.Static("assets"))

	mapRoutes(server)

	server.Run()
}
