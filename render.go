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

// renderer represents a renderer.
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
	t, err := r.g.ParseFile(name + gold.Extension)

	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)

	// Parse the HTML template.
	if err := t.Execute(buf, binding); err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write out the result.
	r.Header().Set(render.ContentType, r.opt.HTMLContentType+r.compiledCharset)
	r.WriteHeader(status)
	io.Copy(r, buf)
}

func (r *renderer) Data(status int, v []byte) {
	r.WriteHeader(status)
}

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

// Renderer is a Middleware that maps a render.Render service into the Martini handler chain. An single variadic rendergold.Options
// struct can be optionally provided to configure HTML rendering. The default directory for templates is "templates" and the default
// file extension is ".gold".
//
// If MARTINI_ENV is set to "" or "development" then templates will be parsed every request. For more performance, set the
// MARTINI_ENV environment variable to "production"
func Renderer(options ...Options) martini.Handler {
	opt := retrieveOptions(options)
	cache := (martini.Env != martini.Dev)
	g := gold.NewGenerator(cache).SetBaseDir(opt.Directory).SetAsset(opt.Asset)
	if opt.Func != nil {
		g.SetHelpers(opt.Func)
	}
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		c.MapTo(&renderer{res, req, opt, compiledCharset(opt), g}, (*render.Render)(nil))
	}
}
