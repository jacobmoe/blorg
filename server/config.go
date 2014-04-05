package server

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Run() {
	server := martini.Classic()

	// setup renderer
	server.Use(render.Renderer(render.Options{
		Layout:     "layouts/application",
		IndentJSON: true,
	}))

	mapRoutes(server)

	server.Run()
}
