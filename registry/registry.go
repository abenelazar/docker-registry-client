package registry

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-logr/logr"
)

/*
 * Pass log messages along to Go's "log" module.
 */
func Log(format string, args ...interface{}) {
	log.Printf(format, args...)
}

type Registry struct {
	URL    string
	Client *http.Client
	log    logr.Logger
	opts   *Options
}

type Options struct {
	Log         logr.Logger
	DisablePing bool
	LogLevel    int
}

/*
 * Create a new Registry with the given URL and credentials, then Ping()s it
 * before returning it to verify that the registry is available.
 *
 * You can, alternately, construct a Registry manually by populating the fields.
 * This passes http.DefaultTransport to WrapTransport when creating the
 * http.Client.
 */
func New(registryURL, username, password string, opts *Options) (*Registry, error) {
	transport := http.DefaultTransport
	return newFromTransport(registryURL, username, password, transport, opts)
}

/*
 * Create a new Registry, as with New, using an http.Transport that disables
 * SSL certificate verification.
 */
func NewInsecure(registryURL, username, password string, opts *Options) (*Registry, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			// TODO: Why?
			InsecureSkipVerify: true, //nolint:gosec
		},
	}

	return newFromTransport(registryURL, username, password, transport, opts)
}

/*
 * Given an existing http.RoundTripper such as http.DefaultTransport, build the
 * transport stack necessary to authenticate to the Docker registry API. This
 * adds in support for OAuth bearer tokens and HTTP Basic auth, and sets up
 * error handling this library relies on.
 */
func WrapTransport(transport http.RoundTripper, url, username, password string) http.RoundTripper {
	tokenTransport := &TokenTransport{
		Transport: transport,
		Username:  username,
		Password:  password,
	}
	basicAuthTransport := &BasicTransport{
		Transport: tokenTransport,
		URL:       url,
		Username:  username,
		Password:  password,
	}
	errorTransport := &ErrorTransport{
		Transport: basicAuthTransport,
	}
	return errorTransport
}

func newFromTransport(registryURL, username, password string, transport http.RoundTripper, opts *Options) (*Registry, error) {
	url := strings.TrimSuffix(registryURL, "/")
	transport = WrapTransport(transport, url, username, password)

	registry := &Registry{
		URL: url,
		Client: &http.Client{
			Transport: transport,
		},
		log:  opts.Log.V(opts.LogLevel),
		opts: opts,
	}

	if opts.DisablePing {
		return registry, nil
	}

	err := registry.Ping()
	return registry, err
}

func (r *Registry) url(pathTemplate string, args ...interface{}) string {
	pathSuffix := fmt.Sprintf(pathTemplate, args...)
	url := fmt.Sprintf("%s%s", r.URL, pathSuffix)
	return url
}

func (r *Registry) Ping() error {
	url := r.url("/v2/")
	r.log.Info("registry.ping url=%s", url)
	resp, err := r.Client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}
