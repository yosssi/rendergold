package rendergold

import (
	"html/template"
	"testing"

	"github.com/martini-contrib/render"
)

func TestRetrieveOptions(t *testing.T) {
	// Case when options == [].
	var options []Options
	opt := retrieveOptions(options)
	if opt.Directory != defaultDirectory ||
		opt.Func != nil ||
		opt.Charset != defaultCharset ||
		opt.HTMLContentType != render.ContentHTML {
		t.Errorf("returned value is invalid [actual: %+v][expected: %+v]",
			opt,
			Options{
				defaultDirectory,
				nil,
				defaultCharset,
				render.ContentHTML,
			},
		)
	}

	// Case when options != [].
	expected := Options{
		"testDir",
		template.FuncMap{"testFunc": func() string { return "test" }},
		"UTF-16",
		"text/xml",
	}
	actual := retrieveOptions([]Options{expected})
	if actual.Directory != expected.Directory ||
		actual.Func == nil ||
		actual.Charset != expected.Charset ||
		actual.HTMLContentType != expected.HTMLContentType {
		t.Errorf("returned value is invalid [actual: %+v][expected: %+v]",
			actual,
			expected,
		)
	}
}
