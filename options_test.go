package rendergold

import (
	"html/template"
	"testing"

	"github.com/martini-contrib/render"
)

func TestOptionsRetrieveFuncMap(t *testing.T) {
	// Case when opt.Funcs == nil.
	opt := Options{}
	actual := opt.retrieveFuncMap()
	if actual != nil {
		t.Errorf("returned value is invalid. [actual: %d][expected: nil]", actual)
	}

	// Case when opt.Funcs != nil.
	expected := template.FuncMap{
		"testFunc": func() string { return "test" },
	}
	opt = Options{Funcs: []template.FuncMap{expected}}
	actual = opt.retrieveFuncMap()
	if actual["testFunc"] == nil {
		t.Errorf("returned value is invalid. [actual: %+v][expected: %+v]", actual, expected)
	}
}

func TestRetrieveOptions(t *testing.T) {
	// Case when options == [].
	var options []Options
	opt := retrieveOptions(options)
	if opt.Directory != defaultDirectory ||
		len(opt.Funcs) != 0 ||
		opt.Charset != defaultCharset ||
		opt.HTMLContentType != render.ContentHTML {
		t.Errorf("returned value is invalid. [actual: %+v][expected: %+v]",
			opt,
			Options{
				defaultDirectory,
				[]template.FuncMap{},
				defaultCharset,
				render.ContentHTML,
			},
		)
	}

	// Case when options != [].
	expected := Options{
		"testDir",
		[]template.FuncMap{
			template.FuncMap{"testFunc": func() string { return "test" }},
		},
		"UTF-16",
		"text/xml",
	}
	actual := retrieveOptions([]Options{expected})
	if actual.Directory != expected.Directory ||
		len(actual.Funcs) != len(expected.Funcs) ||
		actual.Charset != expected.Charset ||
		actual.HTMLContentType != expected.HTMLContentType {
		t.Errorf("returned value is invalid. [actual: %+v][expected: %+v]",
			actual,
			expected,
		)
	}

}
