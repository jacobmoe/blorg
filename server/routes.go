package server

import (
	"github.com/go-martini/martini"
	"github.com/jacobmoe/blorg/controller"
	"github.com/jacobmoe/blorg/controller/api"
)

func mapRoutes(server *martini.ClassicMartini) {

	server.Get("/", controller.Home)

	server.Group("/api", func(apiRoutes martini.Router) {
		apiRoutes.Get("/pages", api.PageIndex)
		apiRoutes.Get("/pages/:id", api.PageShow)
		apiRoutes.Group("/pages/:pageId", func(pageRoutes martini.Router) {
			pageRoutes.Get("/posts", api.PostIndex)
		})
	})
}
