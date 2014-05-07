package rendergold

import "fmt"

// compiledCharset generates a compiled charset and returns it.
func compiledCharset(opt Options) string {
	var cs string
	if opt.Charset == "" {
		cs = defaultCharset
	} else {
		cs = opt.Charset
	}
	return fmt.Sprintf(baseCompiledCharset, cs)
}
