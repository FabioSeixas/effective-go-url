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

// Parse a raw string into a URL.
func Parse(raw string) (*URL, error) {
	index := strings.Index(raw, "://")
	if index < 0 {
		return nil, errors.New("invalid Scheme")
	}
	scheme, rest := raw[:index], raw[index+3:]

	host, path := rest, ""
	index = strings.Index(rest, "/")
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
