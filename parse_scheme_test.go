package url_test

import (
	url "github.com/fabioseixas/effective-go-url"
	"testing"
)

func TestParseScheme(t *testing.T) {
	const (
		rawurl     = "https://foo.com/go"
		wantScheme = "https"
		wantRest   = "foo.com/go"
		wantOk     = true
	)

	scheme, rest, ok := url.ParseScheme(rawurl)
	if scheme != wantScheme {
		t.Errorf("\nparseScheme(%q) \ngot: %q \nwant: %q", rawurl, scheme, wantScheme)
	}

	if rest != wantRest {
		t.Errorf("\nparseScheme(%q) \ngot: %q \nwant: %q", rawurl, rest, wantRest)
	}

	if ok != wantOk {
		t.Errorf("\nparseScheme(%q) \ngot: %v \nwant: %v", rawurl, ok, wantOk)
	}

}
