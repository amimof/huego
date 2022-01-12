package huego

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Client is a simple type used to compose inidivudal requests to an HTTP server.
type Client struct {
	req          *http.Request
	res          *http.Response
	url          *url.URL
	query        string
	apiVersion   string
	verb         string
	username     string
	resourceType string
	resourceName string
	headers      http.Header
	body         io.Reader
	interf       interface{}
	transport    http.RoundTripper
	err          error
}

// Response represents an API response returned by a bridge
type Response struct {
	Data    json.RawMessage `json:"data"`
	Errors  []APIError      `json:"errors,omitempty"`
	DataRaw []byte
	err     error
}

// APIError represents an error returned in a response from a bridge
type APIError struct {
	Description string `json:"description"`
}

// Error satisfies error interface
func (a *APIError) Error() string {
	return a.Description
}

// Get method sets the method on a request to GET. Get will invoke Method(http.MethodGet).
func Get(u string) *Client {
	return newRequest(u).Method(http.MethodGet)
}

// Post method sets the method on a request to POST. Post will invoke Method(http.MethodPost).
func Post(u string) *Client {
	return newRequest(u).Method(http.MethodPost)
}

// Put method sets the method on a request to PUT. Put will invoke Method(http.MethodPut).
func Put(u string) *Client {
	return newRequest(u).Method(http.MethodPut)
}

// Delete method sets the method on a request to DELETE. Delete will invoke Method(http.MethodDelete).
func Delete(u string) *Client {
	return newRequest(u).Method(http.MethodDelete)
}

// Options method sets the method on a request to OPTIONS. Options will invoke Method(http.MethodOptions),
func Options(u string) *Client {
	return newRequest(u).Method(http.MethodOptions)
}

// Transport allows for setting custom roundtripper interfaces used to make http requests
func (r *Client) Transport(t http.RoundTripper) *Client {
	r.transport = t
	return r
}

// Username sets the hue-application-key header for authenticating with a v2 bridge
func (r *Client) Username(u string) *Client {
	r.username = u
	r.Header("hue-application-key", u)
	return r
}

// Method methdo sets the method on a request.
func (r *Client) Method(m string) *Client {
	r.verb = m
	return r
}

// Resource sets the Kubernetes resource to be used when building the URI. For example
// setting the resource to 'Pod' will create an uri like /api/v1/namespaces/pods.
func (r *Client) Resource(res string) *Client {
	r.resourceType = res
	return r
}

// Name sets the name of the Kubernetes resource to be used when building the URI. For example
// setting the name to 'app-pod-1' will create an uri like /api/v1/namespaces/pods/app-pod-1.
func (r *Client) Name(n string) *Client {
	r.resourceName = n
	return r
}

// APIVer sets the api version to be used when building the URI for the request.
// Defaults to 'v1' if not set.
func (r *Client) APIVer(v string) *Client {
	r.apiVersion = v
	return r
}

// Path sets the raw URI path later used by the request.
func (r *Client) Path(p string) *Client {
	r.url.Path = p
	return r
}

// Query sets the raw query path to be used when performing the request
func (r *Client) Query(q string) *Client {
	r.url.RawQuery = q
	return r
}

// Body sets the request body of the request beeing made.
func (r *Client) Body(b io.Reader) *Client {
	r.body = b
	return r
}

// Headers overrides the entire headers field of the http request.
// Use Header() method to set individual headers.
func (r *Client) Headers(h http.Header) *Client {
	r.headers = h
	return r
}

// Header sets one header and replacing any headers with equal key
func (r *Client) Header(key string, values ...string) *Client {
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
func (r *Client) URL() *url.URL {

	// Base path
	p := fmt.Sprintf("/clip/%s/", r.apiVersion)
	if r.url.Path != "" {
		p = r.url.Path
	}

	// Append resource scope
	if len(r.resourceType) != 0 {
		p = path.Join(p, "resource", strings.ToLower(r.resourceType))
	}

	// Append resource name scope
	if len(r.resourceName) != 0 {
		p = path.Join(p, r.resourceName)
	}

	r.url.Path = p

	return r.url
}

// Do executes the request and returns a Response. It uses DoRaw and unmarshals the result into a response type
func (r *Client) Do(ctx context.Context) (*Response, error) {
	d, err := r.DoRaw(ctx)
	if err != nil {
		return nil, err
	}
	var response *Response
	err = json.Unmarshal(d, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// DoRaw executes the request and returns the body of the response
func (r *Client) DoRaw(ctx context.Context) ([]byte, error) {
	// Return any error if any has been generated along the way before continuing
	if r.err != nil {
		return nil, r.err
	}

	u := r.URL().String()
	req, err := http.NewRequestWithContext(ctx, r.verb, u, r.body)
	if err != nil {
		return nil, err
	}
	r.req = req

	if r.headers != nil {
		req.Header = r.headers
	}

	// Make the call
	res, err := r.transport.RoundTrip(r.req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	r.res = res

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Response returns the clients http response. The response will be overwritten for each subsequent request
func (r *Client) Response() *http.Response {
	return r.res
}

// Into sets the interface in which the returning data will be marshaled into.
func (r *Response) Into(obj interface{}) error {
	return json.Unmarshal(r.Data, obj)
}

// GetLightsContext returns an array of light resources
func (r *Client) GetLightsContext(ctx context.Context) ([]*Light, error) {
	res, err := Get(r.url.String()).
		Resource("light").
		Username(r.username).
		Do(ctx)
	if err != nil {
		fmt.Printf("%s\n", "There was an error")
		return nil, err
	}

	var lights []*Light
	err = json.Unmarshal(res.Data, &lights)
	if err != nil {
		fmt.Printf("%s\n", "There was an unmarshaling error")
		return nil, err
	}

	return lights, nil
}

// newRequest returns a Client instance for interaction with a bridge. Hue API V2 and https is used by default.
func newRequest(host string) *Client {
	r := &Client{
		apiVersion: "v2",
		transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	if host == "" {
		r.err = fmt.Errorf("host must be a URL or a host:port pair")
	}

	base := host
	hostURL, err := url.Parse(base)
	if err != nil || hostURL.Scheme == "" || hostURL.Host == "" {
		scheme := "https://"
		hostURL, err = url.Parse(fmt.Sprintf("%s%s", scheme, base))
		if err != nil {
			r.err = err
		}
	}
	r.url = hostURL
	return r
}

// New returns a client using the provided host and username of a bridge
func New(h, u string) *Client {
	return newRequest(h).Username(u)
}
