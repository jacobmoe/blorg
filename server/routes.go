package server

import (
	"github.com/codegangsta/martini"
	"github.com/jacobmoe/blorg/controller"
)

func Run() {
	server := martini.Classic()

	server.Get("/blog", func() string {
		return controller.Blog()
	})

	server.Get("/names/:name", func(args martini.Params) string {
		return controller.Name(args)
	})

	server.Run()
}
