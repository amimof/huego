package huego

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var (
	testTransport = &mockTransport{}
)

func init() {
	SetTransport(testTransport)
}

// mockTransport is a http transport used for mocking http requests
type mockTransport struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

// RoundTrip satisfies http.RoundTripper interface. It's used for mocking http requests
func (c mockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return c.DoFunc(r)
}

func TestSetTransport(t *testing.T) {
	testTransport.DoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			Body:       io.NopCloser(strings.NewReader(`mock transport`)),
			StatusCode: 200,
		}, nil
	}
	SetTransport(testTransport)
	if !reflect.DeepEqual(transport, testTransport) {
		t.Fatalf("want: %+v\n, got: %+v", transport, testTransport)
	}
}
