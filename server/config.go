package server

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Run() {
	server := martini.Classic()

	// setup renderer middleware
	server.Use(render.Renderer(render.Options{
		Layout:     "layouts/application",
		IndentJSON: true,
	}))

	// serve static assets from "assets" in addition to "public"
	server.Use(martini.Static("assets"))

	mapRoutes(server)

	server.Run()
}
