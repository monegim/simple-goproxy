package proxy

import "net/http"

type Options struct {
	Debug    int
	Addr     string
	Upstream string
}

type Proxy struct {
	Opts    *Options
	Version string

	client *http.Client
	server *http.Server
}
const VERSION = "1.6.0"
func NewProxy(opts *Options) (*Proxy, error) {
	proxy := &Proxy{
		Opts: opts,
		Version: VERSION,
	}

	return proxy, nil
}
