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

func parseScheme(rawurl string) (scheme, rest string, ok bool) {
	i := strings.Index(rawurl, "://")
	if i < 1 {
		return "", "", false
	}
	return rawurl[:i], rawurl[i+3:], true
}

// Parse a raw string into a URL.
func Parse(raw string) (*URL, error) {
	scheme, rest, ok := parseScheme(raw)
	if !ok {
		return nil, errors.New("missing scheme")
	}
	host, path := rest, ""
	index := strings.Index(rest, "/")
	if index > 0 {
		host, path = rest[:index], rest[index+1:]
	}

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
