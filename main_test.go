package url

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	rawUrl := "https://foo.com/abc"

	u, err := Parse(rawUrl)

	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawUrl, err)
	}

	want := "https"

	if got := u.Scheme; got != want {
		t.Errorf("Parse(%q).Scheme = %q; want %q", rawUrl, got, want)
	}

	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawUrl, got, want)
	}

	if got, want := u.Path, "abc"; got != want {
		t.Errorf("Parse(%q).Path = %q; want %q", rawUrl, got, want)
	}
}

func TestParseWithoutPath(t *testing.T) {
	rawUrl := "https://foo.com"

	u, err := Parse(rawUrl)

	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawUrl, err)
	}

	want := "https"

	if got := u.Scheme; got != want {
		t.Errorf("Parse(%q).Scheme = %q; want %q", rawUrl, got, want)
	}

	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawUrl, got, want)
	}

	if got, want := u.Path, ""; got != want {
		t.Errorf("Parse(%q).Path = %q; want %q", rawUrl, got, want)
	}
}

func TestURLPort(t *testing.T) {
	tests := []struct {
		in   string
		port string
	}{
		{in: "foo.com:80", port: "80"},
		{in: "foo.com:", port: ""},
		{in: "foo.com", port: ""},
		{in: "1.2.3.4:80", port: "80"},
		{in: "1.2.3.4:", port: ""},
		{in: "1.2.3.4", port: ""},
	}
	for i, test := range tests {
		u := &URL{Host: test.in}

		if got, want := u.Port(), test.port; got != want {
			t.Errorf("test %d - host %q; host.Port() = %q; want %q", i+1, test.in, got, want)
		}
	}

}

func TestURLPortUsingMap(t *testing.T) {
	tests := map[string]struct {
		in   string
		port string
	}{
		"with port":          {in: "foo.com:80", port: "80"},
		"with empty port":    {in: "foo.com:", port: ""},
		"without port":       {in: "foo.com", port: ""},
		"ip with port":       {in: "1.2.3.4:80", port: "80"},
		"ip with empty port": {in: "1.2.3.4:", port: ""},
		"ip without port":    {in: "1.2.3.4", port: ""},
	}
	for name, test := range tests {
		u := &URL{Host: test.in}

		if got, want := u.Port(), test.port; got != want {
			t.Errorf("%q: - host %q; host.Port() = %q; want %q", name, test.in, got, want)
		}
	}

}

func TestURLPortUsingSubtests(t *testing.T) {
	t.Run("with port", func(t *testing.T) {
		const in = "foo.com:80"
		const want = "80"

		u := &URL{Host: in}
		if got := u.Port(); got != want {
			t.Errorf("for host %q; host.Port() = %q; want %q", in, got, want)
		}
	})
	t.Run("without port", func(t *testing.T) {
		const in = "foo.com"
		const want = ""

		u := &URL{Host: in}
		if got := u.Port(); got != want {
			t.Errorf("for host %q; host.Port() = %q; want %q", in, got, want)
		}
	})

}

func TestURLPortUsingSubtestsAndHOF(t *testing.T) {
	testPort := func(in, want string) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			u := &URL{Host: in}

			if got := u.Port(); got != want {
				t.Errorf("for host %q; host.Port() = %q; want %q", in, got, want)
			}
		}
	}

	t.Run("with port", testPort("foo.com:80", "80"))
	t.Run("without port", testPort("foo.com", ""))
}

func TestURLPortUsingSubtestsAndTable(t *testing.T) {
	tests := map[string]struct {
		in   string
		port string
	}{
		"with port":    {in: "foo.com:80", port: "80"},
		"without port": {in: "foo.com", port: ""},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("%s/%s", name, test.in), func(t *testing.T) {

			u := &URL{Host: test.in}

			if got, want := u.Port(), test.port; got != want {
				t.Errorf("for host %q; host.Port() = %q; want %q", test.in, got, want)
			}
		})
	}
}

func TestURLHostname(t *testing.T) {
	tests := map[string]struct {
		in       string
		hostname string
	}{
		"with port":    {in: "foo.com:80", hostname: "foo.com"},
		"without port": {in: "foo.com", hostname: "foo.com"},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("%s/%s", name, test.in), func(t *testing.T) {

			u := &URL{Host: test.in}

			if got, want := u.Hostname(), test.hostname; got != want {
				t.Errorf("for host %q; host.Hostname() = %q; want %q", test.in, got, want)
			}
		})
	}
}

func TestURLString(t *testing.T) {
	u := &URL{
		Scheme: "https",
		Host:   "foo.com",
		Path:   "go",
	}
	got, want := u.String(), "https://foo.com/go"
	if got != want {
		t.Errorf("%#v.String()\ngot %q\nwant %q", u, got, want)
	}
}
