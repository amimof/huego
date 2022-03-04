package huego

import (
	"io"
	"net/http"
)

var (
	testTransport = &mockTransport{}
	testClient    = http.DefaultClient
)

func init() {
	testClient.Transport = testTransport
}

type readCloser struct {
	io.Reader
}

func (readCloser) Close() error {
	return nil
}

func nopCloser(r io.Reader) io.ReadCloser {
	return readCloser{r}
}

// mockTransport is a http transport used for mocking http requests
type mockTransport struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

// RoundTrip satisfies http.RoundTripper interface. It's used for mocking http requests
func (c mockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return c.DoFunc(r)
}

func ptrFloat32(f float32) *float32 {
	return &f
}
func ptrFloat64(f float64) *float64 {
	return &f
}
func ptrString(s string) *string {
	return &s
}
func ptrUint16(u uint16) *uint16 {
	return &u
}
func ptrBool(b bool) *bool {
	return &b
}
