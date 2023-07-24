package huego

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// Client is a simple type used to compose inidivudal requests to an HTTP server.
type CLIPClient struct {
	Client   *http.Client
	baseURL  *url.URL
	username string
}

// Request allows for building a http request
type Request struct {
	c *CLIPClient

	// Generic components
	verb    string
	path    string
	query   string
	headers http.Header

	// Clip components
	apiVersion   string
	resourceType string
	resourceID   string

	// Output
	body io.Reader
	err  error

	// Errorhandler
	onErr func(*http.Response) error
}

// Response represents an API response returned by a bridge
type Response struct {
	Response *http.Response
	BodyRaw  []byte
}

// OnError handler
func (r *Request) OnError(f func(*http.Response) error) *Request {
	r.onErr = f
	return r
}

// Username sets the hue-application-key header for authenticating with a v2 bridge
func (r *Request) Username(u string) *Request {
	if r.c != nil {
		r.c.username = u
		r.Header("hue-application-key", u)
	}
	return r
}

// Verb sets the method on a request
func (r *Request) Verb(m string) *Request {
	r.verb = m
	return r
}

// Resource sets the Kubernetes resource to be used when building the URI. For example
// setting the resource to 'Pod' will create an uri like /api/v1/namespaces/pods.
func (r *Request) Resource(res string) *Request {
	r.resourceType = res
	return r
}

// ID sets the id of the resource to be used when building the URI
func (r *Request) ID(id string) *Request {
	r.resourceID = id
	return r
}

// APIVer sets the api version to be used when building the URI for the request.
// Defaults to 'v1' if not set.
func (r *Request) APIVer(v string) *Request {
	r.apiVersion = v
	return r
}

// Path sets the raw URI path later used by the request.
func (r *Request) Path(p string) *Request {
	r.path = p
	return r
}

// Query sets the raw query path to be used when performing the request
func (r *Request) Query(q string) *Request {
	r.query = q
	return r
}

// Body sets the request body of the request being made.
func (r *Request) Body(data []byte) *Request {
	r.body = bytes.NewReader(data)
	return r
}

// Headers overrides the entire headers field of the http request.
// Use Header() method to set individual headers.
func (r *Request) Headers(h http.Header) *Request {
	r.headers = h
	return r
}

// Header sets one header and replacing any headers with equal key
func (r *Request) Header(key string, values ...string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Del(key)
	for _, value := range values {
		r.headers.Add(key, value)
	}
	return r
}

// URL composes a complete URL and return an url.URL then used by the request
func (r *Request) URL() *url.URL {

	// Base path with api version, default to v2
	if r.apiVersion == "" {
		r.apiVersion = "v2"
	}
	p := fmt.Sprintf("/clip/%s/", r.apiVersion)

	// Append resource scope
	if r.resourceType != "" {
		p = path.Join(p, "resource", strings.ToLower(r.resourceType))
	}

	// Append resource name scope
	if r.resourceID != "" {
		p = path.Join(p, r.resourceID)
	}

	// Use path variable and override everything else
	if r.path != "" {
		p = r.path
	}

	// Parse query params
	q, err := url.ParseQuery(r.query)
	if err != nil {
		r.err = err
	}

	finalURL := &url.URL{}
	if r.c.baseURL != nil {
		*finalURL = *r.c.baseURL
	}
	finalURL.Path = p
	finalURL.RawQuery = q.Encode()

	return finalURL
}

// DoRaw executes the request and returns the body of the response
func (r *Request) Do(ctx context.Context) (*Response, error) {
	// Return any error if any has been generated along the way before continuing
	if r.err != nil {
		return nil, r.err
	}

	client := r.c.Client
	if client == nil {
		client = http.DefaultClient
	}

	u := r.URL().String()
	req, err := http.NewRequestWithContext(ctx, r.verb, u, r.body)
	if err != nil {
		return nil, err
	}

	// Make sure we add auth header
	if r.c.username != "" {
		r.Username(r.c.username)
	}

	if r.headers != nil {
		req.Header = r.headers
	}

	// Make the call
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Handle errors
	if r.onErr == nil {
		r.onErr = func(res *http.Response) error {
			return nil
		}
	}
	err = r.onErr(res)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Response: res,
		BodyRaw:  body,
	}, nil
}

// Into sets the interface in which the returning data will be marshaled into.
func (r *Response) Into(obj interface{}) error {
	err := json.Unmarshal(r.BodyRaw, obj)
	if err != nil {
		return err
	}
	return nil
}

// SetClient is used to set the http.Client to use for making http requests
func (c *CLIPClient) SetClient(client *http.Client) *CLIPClient {
	c.Client = client
	return c
}

// NewClient creates a client for making http requests
func NewClient(h, u string) (*CLIPClient, error) {
	if h == "" {
		return nil, fmt.Errorf("host must be a URL or a host:port pair")
	}
	base := h
	hostURL, err := url.Parse(base)
	if err != nil || hostURL.Scheme == "" || hostURL.Host == "" {
		scheme := "https://"
		hostURL, err = url.Parse(fmt.Sprintf("%s%s", scheme, base))
		if err != nil {
			return nil, err
		}
	}
	return &CLIPClient{
		Client:   http.DefaultClient,
		username: u,
		baseURL:  hostURL,
	}, nil
}

// NewInsecureClient creates an insecure client for making http requests.
// It sets InsecureSkipVerify to true on the underlying Transport
func NewInsecureClient(h, u string) (*CLIPClient, error) {
	c, err := NewClient(h, u)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	c.Client = client
	return c, nil
}

// NewRequest creates a default request using the given client
func NewRequest(c *CLIPClient) *Request {
	return &Request{
		c:          c,
		apiVersion: "v2",
	}
}

// NewClientV2 creates a client for V2 bridges
func NewClientV2(h, u string) (*ClientV2, error) {
	cc, err := NewClient(h, u)
	if err != nil {
		return nil, err
	}
	return &ClientV2{
		Clip: cc,
	}, nil
}

// NewInsecureClientV2 creates an insecure client for V2 bridges
func NewInsecureClientV2(h, u string) (*ClientV2, error) {
	cc, err := NewInsecureClient(h, u)
	if err != nil {
		return nil, err
	}
	return &ClientV2{
		Clip: cc,
	}, nil
}
