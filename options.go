package rendergold

import (
	"html/template"

	"github.com/martini-contrib/render"
)

// Options is a struct for specifying configuration options for the render.Renderer middleware
type Options struct {
	// Directory to load templates. Default is "templates"
	Directory string
	// Funcs is a slice of FuncMaps to apply to the template upon compilation. This is useful for helper functions. Defaults to [].
	Func template.FuncMap
	// Appends the given charset to the Content-Type header. Default is "UTF-8".
	Charset string
	// Allows changing of output to XHTML instead of HTML. Default is "text/html"
	HTMLContentType string
}

// retrieveOptions retrieves an options from the array of options
// and sets default values to it.
func retrieveOptions(options []Options) Options {
	var opt Options

	if len(options) > 0 {
		opt = options[0]
	}

	// Set a defalut directory to the option.
	if opt.Directory == "" {
		opt.Directory = defaultDirectory
	}

	// Set a defalut char set to the option.
	if opt.Charset == "" {
		opt.Charset = defaultCharset
	}

	// Set a defalut HTML content type to the option.
	if opt.HTMLContentType == "" {
		opt.HTMLContentType = render.ContentHTML
	}

	return opt
}
