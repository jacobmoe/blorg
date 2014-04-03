package main

import (
	"github.com/codegangsta/martini"
	"github.com/jacobmoe/gorg"
	"io/ioutil"
	"path/filepath"
)

func main() {

	inPath, _ := filepath.Abs("test/test.org")
	outPath, _ := filepath.Abs("test/test.html")

	gorg.OrgToHtmlFile(inPath, outPath)

	server := martini.Classic()
	server.Get("/", func() string {
		return ioutil.ReadFile(outPath)
	})

	server.Run()
}
