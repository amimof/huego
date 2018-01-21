// Package huego provides an extensive, easy to use interface to the Philips Hue bridge.
package huego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

// Response is a wrapper struct of the success response returned from the bridge after a successfull API call.
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

func get(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
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

func put(url string, data []byte) ([]byte, error) {

	body := strings.NewReader(string(data))

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
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

func post(url string, data []byte) ([]byte, error) {

	body := strings.NewReader(string(data))

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
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

func delete(url string) ([]byte, error) {

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
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

// DiscoverAll performs a discovery on the network looking for bridges using https://www.meethue.com/api/nupnp service
func DiscoverAll() ([]Bridge, error) {

	res, err := get("https://www.meethue.com/api/nupnp")
	if err != nil {
		return nil, err
	}

	var bridges []Bridge

	err = json.Unmarshal(res, &bridges)
	if err != nil {
		return nil, err
	}

	return bridges, nil

}

// Discover performs a discovery on the network looking for bridges using https://www.meethue.com/api/nupnp service. Returns the first bridge if any found
func Discover() (*Bridge, error) {

	var b *Bridge

	bridges, err := DiscoverAll()
	if err != nil {
		return nil, err
	}

	if len(bridges) > 0 {
		b = &bridges[0]
	}

	return b, nil

}

// New instantiates and returns a new Bridge
func New(h, u string) *Bridge {
	return &Bridge{
		Host: h,
		User: u,
		ID:   "",
	}
}
