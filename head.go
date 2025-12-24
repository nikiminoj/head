package head

import (
	"context"
	"net/http"
)

type Config struct{}

func CreateConfig() *Config {
	return &Config{}
}

type Middleware struct {
	next http.Handler
}

func New(
	ctx context.Context,
	next http.Handler,
	config *Config,
	name string,
) (http.Handler, error) {
	return &Middleware{next: next}, nil
}

func (m *Middleware) ServeHTTP(
	rw http.ResponseWriter,
	req *http.Request,
) {
	// === NGINX equivalent ===
	// proxy_set_header X-Original-URI $request_uri;
	req.Header.Set(
		"X-Original-URI",
		req.URL.RequestURI(),
	)

	m.next.ServeHTTP(rw, req)
}
