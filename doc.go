// Package rendergold is a middleware for Martini that provides Gold templates parsing and HTML rendering.
//
//  package main
//
//  import (
//    "github.com/go-martini/martini"
//    "github.com/yosssi/rendergold"
//  )
//
//  func main() {
//    m := martini.Classic()
//    m.Use(rendergold.Renderer()) // reads "templates" directory by default
//
//    m.Get("/html", func(r render.Render) {
//      r.HTML(200, "mytemplate", nil)
//    })
//
//    m.Get("/json", func(r render.Render) {
//      r.JSON(200, "hello world")
//    })
//
//    m.Run()
//  }
package rendergold
