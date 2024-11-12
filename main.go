package url

import (
	"errors"
	"fmt"
	"strings"
)

// Represents an URL
type URL struct {
	Scheme string
	Host   string
	Path   string
}

func parseScheme(raw string) (scheme, rest string, ok bool) {
	i := strings.Index(raw, "://")
	if i < 1 {
		return "", "", false
	}
	return raw[:i], raw[i+3:], true
}

func parseHostAndPath(raw string) (host, path string) {
	host, path = raw, ""
	index := strings.Index(raw, "/")
	if index > 0 {
		return raw[:index], raw[index+1:]
	}
	return host, path
}

// Parse a raw string into a URL.
func Parse(raw string) (*URL, error) {
	scheme, rest, ok := parseScheme(raw)
	if !ok {
		return nil, errors.New("missing scheme")
	}
	host, path := parseHostAndPath(rest)

	return &URL{Scheme: scheme, Host: host, Path: path}, nil
}

func (url *URL) Port() string {
	port := ""
	i := strings.Index(url.Host, ":")

	if i > 0 {
		port = url.Host[i+1:]
	}

	return port
}

func (url *URL) Hostname() string {
	i := strings.Index(url.Host, ":")

	if i > 0 {
		return url.Host[:i]
	}

	return url.Host
}

func (url *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", url.Scheme, url.Host, url.Path)
}
