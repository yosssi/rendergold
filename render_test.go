package rendergold

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/yosssi/gold"
)

func TestRenderHTML(t *testing.T) {
	// Case when r.g.ParseFile returns an error.
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	opt := Options{
		Directory:       "test",
		Func:            nil,
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r := renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.HTML(http.StatusOK, "notExistTemplate", nil)
	if res.Code != http.StatusInternalServerError {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusInternalServerError)
	}

	// Case when t.Execute returns an error.
	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	opt = Options{
		Directory:       "test",
		Func:            nil,
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r = renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.HTML(http.StatusOK, "0002", "Test")
	if res.Code != http.StatusInternalServerError {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusInternalServerError)
	}

	// Case when no error occurs (with no HTMLOptions).
	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	opt = Options{
		Directory:       "test",
		Func:            nil,
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r = renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.HTML(http.StatusOK, "0001", nil)
	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}

	// Case when no error occurs (with HTMLOptions).
	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	opt = Options{
		Directory:       "test",
		Func:            nil,
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r = renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.HTML(http.StatusOK, "0001", nil, render.HTMLOptions{Layout: "0001" + NameContentDelim + "html"})
	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}
}

func TestRenderRedirect(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	opt := Options{
		Directory:       defaultDirectory,
		Func:            nil,
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r := renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.Redirect("https://github.com", http.StatusMovedPermanently)
	if res.Code != http.StatusMovedPermanently {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusMovedPermanently)
	}
}

func TestRenderer(t *testing.T) {
	opt := Options{
		Directory: "test",
		Func: template.FuncMap{
			"testFunc": func() string { return "test" },
		},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	Renderer(opt)

	// Test via a Martini context.
	m := martini.Classic()
	m.Use(Renderer(opt))
	m.Get("/", func(r Render) {
		r.HTML(http.StatusOK, "0001", nil)
	})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}
}
