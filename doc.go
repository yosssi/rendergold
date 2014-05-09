// Package rendergold provides a Martini middleware/handler for parsing Gold templates and rendering HTML.
//
//  package main
//
//  import (
//    "github.com/go-martini/martini"
//    "github.com/martini-contrib/render"
//    "github.com/yosssi/rendergold"
//  )
//
//  func main() {
//    m := martini.Classic()
//    m.Use(rendergold.Renderer()) // reads "templates" directory by default
//
//    m.Get("/", func(r render.Render) {
//      r.HTML(200, "mytemplate", nil)
//    })
//
//    m.Run()
//  }
package rendergold
