package server

import (
	"github.com/go-martini/martini"
	"github.com/jacobmoe/blorg/controller"
	"github.com/martini-contrib/render"
)

func Run() {
	server := martini.Classic()
	server.Use(render.Renderer())

	server.Get("/", controller.Home)
	server.Get("/blog", controller.Blog)
	server.Get("/names/:name", controller.Name)

	server.Run()
}
