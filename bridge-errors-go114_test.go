//+build go1.14

package huego

import (
	"context"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestBridge_getAPIPathError(t *testing.T) {
	b := New("invalid hostname", "")
	expected := "parse \"http://invalid hostname\": invalid character \" \" in host name"
	_, err := b.getAPIPath("/")
	if errString := err.Error(); errString != expected {
		t.Fatalf("Expected error %s but got %s", expected, errString)
	}
}

func TestBridge_getError(t *testing.T) {
	httpmock.Deactivate()
	defer httpmock.Activate()
	expected := "Get \"invalid%20hostname\": unsupported protocol scheme \"\""
	_, err := get(context.Background(), "invalid hostname")
	if errString := err.Error(); errString != expected {
		t.Fatalf("Expected error %s but got %s", expected, errString)
	}
}

func TestBridge_putError(t *testing.T) {
	httpmock.Deactivate()
	defer httpmock.Activate()
	expected := "Put \"invalid%20hostname\": unsupported protocol scheme \"\""
	_, err := put(context.Background(), "invalid hostname", []byte("huego"))
	if errString := err.Error(); errString != expected {
		t.Fatalf("Expected error %s but got %s", expected, errString)
	}
}

func TestBridge_postError(t *testing.T) {
	httpmock.Deactivate()
	defer httpmock.Activate()
	expected := "Post \"invalid%20hostname\": unsupported protocol scheme \"\""
	_, err := post(context.Background(), "invalid hostname", []byte("huego"))
	if errString := err.Error(); errString != expected {
		t.Fatalf("Expected error %s but got %s", expected, errString)
	}
}

func TestBridge_deleteError(t *testing.T) {
	httpmock.Deactivate()
	defer httpmock.Activate()
	expected := "Delete \"invalid%20hostname\": unsupported protocol scheme \"\""
	_, err := delete(context.Background(), "invalid hostname")
	if errString := err.Error(); errString != expected {
		t.Fatalf("Expected error %s but got %s", expected, errString)
	}
}
