// Package huego provides an easy to use interface to the Philips Hue bridge. 
package huego

import (
	"net/http"
	"encoding/json"
	"net/url"
	"path"
	"strings"
	"io/ioutil"
	"fmt"
)

type Bridge struct {
	Host string
	User string
}

type ApiResponse struct {
	Success map[string]interface{} `json:"success,omitempty"`
	Error *ApiError `json:"error,omitempty"`
}

type ApiError struct {
	Type int
	Address string
	Description string
}

type Response struct {
	Success map[string]interface{}
}

func (a *ApiError) UnmarshalJSON(data []byte) error {
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

func (r *ApiError) Error() string {
	return fmt.Sprintf("ERROR %d [%s]: \"%s\"", r.Type, r.Address, r.Description)
}

func handleResponse(a []*ApiResponse) (*Response, error) {
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

func (b *Bridge) getApiPath(str ...string) (string, error) {
	u, err := url.Parse(b.Host)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, "/api/", b.User)
	for _, p := range str {
		u.Path = path.Join(u.Path, p)
	}
	return u.String(), nil
}

func (b *Bridge) GetApiUrl(str ...string) string {
	u, err := url.Parse(b.Host)
	if err != nil {
		return ""
	}
	u.Path = path.Join(u.Path, "/api/", b.User)
	for _, p := range str {
		u.Path = path.Join(u.Path, p)
	}
	return u.String()
}

func (b *Bridge) getResource(url string) ([]byte, error) {

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

func (b *Bridge) putResource(url string, data []byte) ([]byte, error) {

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

func (b *Bridge) postResource(url string, data []byte) ([]byte, error) {

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

func (b *Bridge) deleteResource(url string) ([]byte, error) {

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

// Instantiates and returns a new Bridge object
func New(h, u string) *Bridge {
	return &Bridge{h, u}
}
