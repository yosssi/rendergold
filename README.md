# RenderGold [![wercker status](https://app.wercker.com/status/3cb287c2c7570a2cf024eab2da6d7a14/s/ "wercker status")](https://app.wercker.com/project/bykey/3cb287c2c7570a2cf024eab2da6d7a14) [![GoDoc](https://godoc.org/github.com/yosssi/rendergold?status.png)](https://godoc.org/github.com/yosssi/rendergold)

Martini middleware/handler for parsing [Gold](http://gold.yoss.si/) templates and rendering HTML.

## About Gold

[Gold](http://gold.yoss.si/) is a template engine for Go. Please visit [its GitHub repository](https://github.com/yosssi/gold) for more details.

## Usage

Here is an example of server-side Go code:

```go
package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/yosssi/rendergold"
)

func main() {
	m := martini.Classic()
	m.Use(rendergold.Renderer()) // reads "templates" directory by default

	m.Get("/", func(r render.Render) {
		r.HTML(200, "top", nil) // parses "templates/top.gold"
	})

	m.Run()
}
```

Here is an example of `templates/top.gold`:

```gold
doctype html
html
  head
    title RenderGold
  body
    h1 Hello RenderGold!
```

This template will be converted to the following HTML:

```hmtl
<!DOCTYPE html>
<html>
  <head>
    <title>RenderGold</title>
  </head>
  <body>
    <h1>Hello RenderGold!</h1>
  </body>
</html>
```

## Options

`rendergold.Renderer` comes with a variety of configuration options:

~~~ go
// ...
m.Use(render.Renderer(render.Options{
  Directory: "templates", // Specify what path to load the templates from. Default is "templates".
  Func: template.FuncMap{AppHelpers}, // Specify helper function map for templates to access.
  Charset: "UTF-8", // Sets encoding for html content-types. Default is "UTF-8".
  HTMLContentType: "application/xhtml+xml", // Output XHTML content type instead of default "text/html"
}))
// ...
~~~

## Parse strings instead of template files

RenderGold can parse strings from memory instead of template files on disk:

```go
package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/yosssi/rendergold"
)

func main() {
	m := martini.Classic()
	m.Use(rendergold.Renderer()) // reads "templates" directory by default

	m.Get("/", func(r render.Render) {
		// template strings
		parent := `
doctype html
html
  head
    title Gold
  body
    block content
    footer
      block footer
`
		child := `
extends parent

block content
  #container
    | Hello Gold

block footer
  .footer
    | Copyright XXX
`
		// parses strings instead of template files
		r.HTML(
			200,
			"child",
			nil,
			render.HTMLOptions{
				Layout: "parent" + rendergold.NameContentDelim + parent,
			},
			render.HTMLOptions{
				Layout: "child" + rendergold.NameContentDelim + child,
			},
		)
	})

	m.Run()
}
```

## Docs

* [GoDoc](https://godoc.org/github.com/yosssi/rendergold)
