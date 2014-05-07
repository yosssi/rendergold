// Package rendergold is a middleware for Martini that provides Gold template rendering.
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

import (
	"bytes"
	"html/template"
	"io"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/yosssi/gold"
)

// Options is a struct for specifying configuration options for the render.Renderer middleware
type Options struct {
	// Directory to load templates. Default is "templates"
	Directory string
	// Funcs is a slice of FuncMaps to apply to the template upon compilation. This is useful for helper functions. Defaults to [].
	Funcs []template.FuncMap
	// Appends the given charset to the Content-Type header. Default is "UTF-8".
	Charset string
	// Allows changing of output to XHTML instead of HTML. Default is "text/html"
	HTMLContentType string
}

// Renderer is a Middleware that maps a render.Render service into the Martini handler chain. An single variadic rendergold.Options
// struct can be optionally provided to configure HTML rendering. The default directory for templates is "templates" and the default
// file extension is ".gold".
//
// If MARTINI_ENV is set to "" or "development" then templates will be parsed every request. For more performance, set the
// MARTINI_ENV environment variable to "production"
func Renderer(options ...Options) martini.Handler {
	opt := prepareOptions(options)
	cs := prepareCharset(opt.Charset)
	var cache bool
	if martini.Env == martini.Dev {
		// parse templates every request
		cache = false
	} else {
		// cache parsed results
		cache = true
	}
	g := gold.NewGenerator(cache)
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		c.MapTo(&renderer{res, req, opt, cs, g}, (*render.Render)(nil))
	}
}

func prepareOptions(options []Options) Options {
	var opt Options
	if len(options) > 0 {
		opt = options[0]
	}

	// Defaults
	if len(opt.Directory) == 0 {
		opt.Directory = "templates"
	}
	if len(opt.HTMLContentType) == 0 {
		opt.HTMLContentType = render.ContentHTML
	}

	return opt
}

func prepareCharset(charset string) string {
	if len(charset) != 0 {
		return "; charset=" + charset
	}

	return "; charset=UTF-8"
}

type renderer struct {
	http.ResponseWriter
	req             *http.Request
	opt             Options
	compiledCharset string
	g               *gold.Generator
}

func (r *renderer) JSON(status int, v interface{}) {
	r.WriteHeader(status)
}

func (r *renderer) HTML(status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) {
	out, err := r.execute(name, binding)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	// template rendered fine, write out the result
	r.Header().Set(render.ContentType, r.opt.HTMLContentType+r.compiledCharset)
	r.WriteHeader(status)
	io.Copy(r, out)
}

func (r *renderer) Data(status int, v []byte) {
	r.WriteHeader(status)
}

// Error writes the given HTTP status to the current ResponseWriter
func (r *renderer) Error(status int) {
	r.WriteHeader(status)
}

func (r *renderer) Status(status int) {
	r.WriteHeader(status)
}

func (r *renderer) Redirect(location string, status ...int) {
	code := http.StatusFound
	if len(status) == 1 {
		code = status[0]
	}

	http.Redirect(r, r.req, location, code)
}

func (r *renderer) Template() *template.Template {
	return nil
}

func (r *renderer) execute(name string, binding interface{}) (*bytes.Buffer, error) {
	t, err := r.g.ParseFile(r.opt.Directory + "/" + name + ".gold")
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	return buf, t.Execute(buf, binding)
}
