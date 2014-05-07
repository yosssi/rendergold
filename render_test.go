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

func TestRenderJSON(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	opt := Options{
		Directory:       defaultDirectory,
		Funcs:           []template.FuncMap{},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r := renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.JSON(http.StatusOK, nil)
	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}
}

func TestRenderHTML(t *testing.T) {
	// Case when r.g.ParseFile returns an error.
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	opt := Options{
		Directory:       "test",
		Funcs:           []template.FuncMap{},
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
		Funcs:           []template.FuncMap{},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r = renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.HTML(http.StatusOK, "0002", "Test")
	if res.Code != http.StatusInternalServerError {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusInternalServerError)
	}

	// Case when no error occurs.
	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	opt = Options{
		Directory:       "test",
		Funcs:           []template.FuncMap{},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r = renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.HTML(http.StatusOK, "0001", nil)
	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}
}

func TestRenderData(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	opt := Options{
		Directory:       defaultDirectory,
		Funcs:           []template.FuncMap{},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r := renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.Data(http.StatusOK, nil)
	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}
}

func TestRenderError(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	opt := Options{
		Directory:       defaultDirectory,
		Funcs:           []template.FuncMap{},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r := renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.Error(http.StatusInternalServerError)
	if res.Code != http.StatusInternalServerError {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusInternalServerError)
	}
}

func TestRenderStatus(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	opt := Options{
		Directory:       defaultDirectory,
		Funcs:           []template.FuncMap{},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r := renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.Status(http.StatusOK)
	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}
}

func TestRenderRedirect(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	opt := Options{
		Directory:       defaultDirectory,
		Funcs:           []template.FuncMap{},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r := renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	r.Redirect("https://github.com", http.StatusMovedPermanently)
	if res.Code != http.StatusMovedPermanently {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusMovedPermanently)
	}
}

func TestRenderTemplate(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	opt := Options{
		Directory:       defaultDirectory,
		Funcs:           []template.FuncMap{},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	r := renderer{res, req, opt, compiledCharset(opt), gold.NewGenerator(false).SetBaseDir(opt.Directory)}
	tpl := r.Template()
	if tpl != nil {
		t.Errorf("returned value should be nil [actual: %+v]", tpl)
	}
}

func TestRenderer(t *testing.T) {
	opt := Options{
		Directory: "test",
		Funcs: []template.FuncMap{
			template.FuncMap{
				"testFunc": func() string { return "test" },
			},
		},
		Charset:         defaultCharset,
		HTMLContentType: render.ContentHTML,
	}
	Renderer(opt)

	// Test via a Martini context.
	m := martini.Classic()
	m.Use(Renderer(opt))
	m.Get("/", func(r render.Render) {
		r.HTML(http.StatusOK, "0001", nil)
	})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}
}
