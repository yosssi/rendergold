package rendergold

import (
	"fmt"
	"testing"
)

func TestCompiledCharset(t *testing.T) {
	// Case when opt.Charset == "".
	opt := Options{}
	actual := compiledCharset(opt)
	expected := fmt.Sprintf(baseCompiledCharset, defaultCharset)
	if actual != expected {
		t.Errorf("returned value is invalid. [actual: %d][expected: %d]", actual, expected)
	}

	// Case when opt.Charset != "".
	charset := "UTF-16"
	opt = Options{Charset: charset}
	actual = compiledCharset(opt)
	expected = fmt.Sprintf(baseCompiledCharset, charset)
	if actual != expected {
		t.Errorf("returned value is invalid. [actual: %d][expected: %d]", actual, expected)
	}
}
