package api

import (
	"github.com/go-martini/martini"
	"github.com/jacobmoe/blorg/model"
	"github.com/jacobmoe/gorg"
	// "github.com/martini-contrib/render"
	"github.com/jacobmoe/render"
	"path/filepath"
	"strconv"
)

func PageIndex(render render.Render) {
	inPath, _ := filepath.Abs("example_files/example.org")

	tree := gorg.TreeFromFile(inPath)

	pages := model.PagesFromTree(tree)

	render.JSON(200, pages)
}

func PageShow(args martini.Params, render render.Render) {
	inPath, _ := filepath.Abs("example_files/example.org")

	tree := gorg.TreeFromFile(inPath)

	pages := model.PagesFromTree(tree)

	i, err := strconv.Atoi(args["id"])

	if i <= 0 || i > len(pages) || err != nil {
		render.JSON(404, "404 nothing here")
	} else {
		render.JSON(200, pages[i-1])
	}
}
