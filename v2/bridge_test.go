package huego

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestDiscover(t *testing.T) {

	tests := []struct {
		Body       string
		StatusCode int
		Error      error
		Expect     *DiscoveredBridge
	}{
		// Return one bridge
		{
			Body:       `[{"id":"001788fffe73ff19","internalipaddress":"192.168.13.112","port":443}]`,
			StatusCode: 200,
			Error:      nil,
			Expect: &DiscoveredBridge{
				InternalIPAddress: "192.168.13.112",
				Port:              443,
				ID:                "001788fffe73ff19",
			},
		},
		// Return more than one bridge but result should be the first one
		{
			Body:       `[{"id":"001788fffe73ff19","internalipaddress":"192.168.13.112","port":443},{"id":"fe73ff19001788ff","internalipaddress":"172.16.1.112","port":8443}]`,
			StatusCode: 200,
			Error:      nil,
			Expect: &DiscoveredBridge{
				InternalIPAddress: "192.168.13.112",
				Port:              443,
				ID:                "001788fffe73ff19",
			},
		},
		// Return 0 bridges which will cause an error
		{
			Body:       `[]`,
			StatusCode: 200,
			Error:      fmt.Errorf("no bridges found during discovery"),
			Expect:     nil,
		},
		// Return an error caused by unexpected response body
		{
			Body:       `unexpected text`,
			StatusCode: 200,
			Error:      fmt.Errorf("invalid character 'u' looking for beginning of value"),
			Expect:     nil,
		},
		// Return an error caused by unexpected json response
		{
			Body:       `{"message":"ok"}`,
			StatusCode: 200,
			Error:      fmt.Errorf("json: cannot unmarshal object into Go value of type []huego.DiscoveredBridge"),
			Expect:     nil,
		},
	}

	for _, test := range tests {
		testTransport.DoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				Body:       io.NopCloser(strings.NewReader(test.Body)),
				StatusCode: test.StatusCode,
			}, nil
		}
		b, err := Discover()
		if err != nil && err.Error() != test.Error.Error() {
			t.Fatalf("\nwant: %+v\n, got: %+v", test.Error, err)
		}
		if !reflect.DeepEqual(b, test.Expect) {
			t.Fatalf("want: %+v\n, got: %+v", test.Expect, b)
		}
	}

}

func TestDiscoverAll(t *testing.T) {
	tests := []struct {
		Body       string
		StatusCode int
		Error      error
		Expect     []DiscoveredBridge
	}{
		// Return one bridge
		{
			Body:       `[{"id":"001788fffe73ff19","internalipaddress":"192.168.13.112","port":443}]`,
			StatusCode: 200,
			Error:      nil,
			Expect: []DiscoveredBridge{
				{
					InternalIPAddress: "192.168.13.112",
					Port:              443,
					ID:                "001788fffe73ff19",
				},
			},
		},
		// Return more than one bridge
		{
			Body:       `[{"id":"001788fffe73ff19","internalipaddress":"192.168.13.112","port":443},{"id":"fe73ff19001788ff","internalipaddress":"172.16.1.112","port":8443}]`,
			StatusCode: 200,
			Error:      nil,
			Expect: []DiscoveredBridge{
				{
					InternalIPAddress: "192.168.13.112",
					Port:              443,
					ID:                "001788fffe73ff19",
				},
				{
					InternalIPAddress: "172.16.1.112",
					Port:              8443,
					ID:                "fe73ff19001788ff",
				},
			},
		},
	}

	for _, test := range tests {
		testTransport.DoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				Body:       io.NopCloser(strings.NewReader(test.Body)),
				StatusCode: test.StatusCode,
			}, nil
		}
		b, err := DiscoverAll()
		if err != nil && err.Error() != test.Error.Error() {
			t.Fatalf("\nwant: %+v\n, got: %+v", test.Error, err)
		}
		if !reflect.DeepEqual(b, test.Expect) {
			t.Fatalf("want: %+v\n, got: %+v", test.Expect, b)
		}
	}

}
