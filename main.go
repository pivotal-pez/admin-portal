package main

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

const (
	SuccessStatus = 200
)

func main() {
	m := martini.Classic()

	m.Get("/", func(params martini.Params, log *log.Logger, r render.Render) {
		r.HTML(SuccessStatus, "index", nil)
	})

	m.Run()
}
