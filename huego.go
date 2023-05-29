// Package huego provides an extensive, easy to use interface to the Philips Hue bridge.
package huego

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	applicationJSON = "application/json"
	contentType     = "Content-Type"
)

// APIResponse holds the response data returned form the bridge after a request has been made.
type APIResponse struct {
	Success map[string]interface{} `json:"success,omitempty"`
	Error   *APIError              `json:"error,omitempty"`
}

// APIError defines the error response object returned from the bridge after an invalid API request.
type APIError struct {
	Type        int
	Address     string
	Description string
}

// Response is a wrapper struct of the success response returned from the bridge after a successful API call.
type Response struct {
	Success map[string]interface{}
}

// UnmarshalJSON makes sure that types are correct when unmarshalling. Implements package encoding/json
func (a *APIError) UnmarshalJSON(data []byte) error {
	var aux map[string]interface{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	a.Type = int(aux["type"].(float64))
	a.Address = aux["address"].(string)
	a.Description = aux["description"].(string)
	return nil
}

// Error returns an error string
func (a *APIError) Error() string {
	return fmt.Sprintf("ERROR %d [%s]: \"%s\"", a.Type, a.Address, a.Description)
}

func handleResponse(a []*APIResponse) (*Response, error) {
	success := map[string]interface{}{}
	for _, r := range a {
		if r.Success != nil {
			for k, v := range r.Success {
				success[k] = v
			}
		}
		if r.Error != nil {
			return nil, r.Error
		}
	}
	resp := &Response{Success: success}
	return resp, nil
}

// unmarshal will try to unmarshal data into APIResponse so that we can
// return the actual error returned by the bridge http API as an error struct.
func unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, &v)
	if err != nil {
		var a []*APIResponse
		err = json.Unmarshal(data, &a)
		if err != nil {
			return err
		}
		_, err = handleResponse(a)
		if err != nil {
			return err
		}
	}
	return nil
}

func get(ctx context.Context, url string, client *http.Client) ([]byte, error) {
	if os.Getenv("HUEGO_DEBUG") != "" {
		fmt.Printf("GET %s\n", url)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func put(ctx context.Context, url string, data []byte, client *http.Client) ([]byte, error) {
	if os.Getenv("HUEGO_DEBUG") != "" {
		fmt.Printf("PUT %s\n", url)
		fmt.Println(string(data))
	}

	body := strings.NewReader(string(data))

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set(contentType, applicationJSON)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func post(ctx context.Context, url string, data []byte, client *http.Client) ([]byte, error) {
	if os.Getenv("HUEGO_DEBUG") != "" {
		fmt.Printf("POST %s\n", url)
		fmt.Println(string(data))
	}

	body := strings.NewReader(string(data))

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set(contentType, applicationJSON)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func del(ctx context.Context, url string, client *http.Client) ([]byte, error) {
	if os.Getenv("HUEGO_DEBUG") != "" {
		fmt.Printf("DEL %s\n", url)
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set(contentType, applicationJSON)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// DiscoverAll performs a discovery on the network looking for bridges using https://www.meethue.com/api/nupnp service.
// DiscoverAll returns a list of Bridge objects.
func DiscoverAll() ([]Bridge, error) {
	return DiscoverAllContext(context.Background())
}

// DiscoverAllContext performs a discovery on the network looking for bridges using https://www.meethue.com/api/nupnp service.
// DiscoverAllContext returns a list of Bridge objects.
func DiscoverAllContext(ctx context.Context) ([]Bridge, error) {

	req, err := http.NewRequest(http.MethodGet, "https://discovery.meethue.com", nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var bridges []Bridge

	err = json.Unmarshal(d, &bridges)
	if err != nil {
		return nil, err
	}

	return bridges, nil

}

// Discover performs a discovery on the network looking for bridges using https://www.meethue.com/api/nupnp service.
// Discover uses DiscoverAll() but only returns the first instance in the array of bridges if any.
func Discover() (*Bridge, error) {
	return DiscoverContext(context.Background())
}

// DiscoverContext performs a discovery on the network looking for bridges using https://www.meethue.com/api/nupnp service.
// DiscoverContext uses DiscoverAllContext() but only returns the first instance in the array of bridges if any.
func DiscoverContext(ctx context.Context) (*Bridge, error) {
	b := &Bridge{}

	bridges, err := DiscoverAllContext(ctx)
	if err != nil {
		return nil, err
	}

	if len(bridges) > 0 {
		b = &bridges[0]
	}

	return b, nil

}

// New instantiates and returns a new Bridge. New accepts hostname/ip address to the bridge (h) as well as an username (u).
// h may or may not be prefixed with http(s)://. For example http://192.168.1.20/ or 192.168.1.20.
// u is a username known to the bridge. Use Discover() and CreateUser() to create a user.
func New(h, u string) *Bridge {
	return &Bridge{
		Host:   h,
		User:   u,
		ID:     "",
		client: http.DefaultClient,
	}
}

/*
NewWithClient instantiates and returns a new Bridge with a custom HTTP client.
NewWithClient accepts the same parameters as New, but with an additional acceptance of an http.Client.

  - h may or may not be prefixed with http(s)://. For example http://192.168.1.20/ or 192.168.1.20.
  - u is a username known to the bridge. Use Discover() and CreateUser() to create a user.
  - Difference between New and NewWithClient being the ability to implement your own http.RoundTripper for proxying.
*/
func NewWithClient(h, u string, client *http.Client) *Bridge {
	return &Bridge{
		Host:   h,
		User:   u,
		ID:     "",
		client: client,
	}
}

/*
NewCustom instantiates and returns a new Bridge. NewCustom accepts:
  - a raw JSON []byte slice as input for substantiating the Bridge type
  - a custom HTTP client like NewWithClient that will be used to make API requests

Note that this is for advanced users, the other "New" functions may suit you better.
*/
func NewCustom(raw []byte, host string, client *http.Client) (*Bridge, error) {
	br := &Bridge{}
	if err := json.Unmarshal(raw, br); err != nil {
		return nil, err
	}
	br.Host = host
	br.client = client
	return br, nil
}
