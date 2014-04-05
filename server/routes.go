package server

import (
	"github.com/go-martini/martini"
	"github.com/jacobmoe/blorg/controller"
	"github.com/jacobmoe/blorg/controller/api"
)

func mapRoutes(server *martini.ClassicMartini) {

	server.Get("/", controller.Home)
	server.Get("/blog", controller.Blog)
	server.Get("/names/:name", controller.Name)

	server.Get("/api/blog", api.BlogShow)
}
