package huego

import (
	//"github.com/amimof/loglevel-go"
	"net/http"
	//"crypto/tls"
	"encoding/json"

	"net/url"

	"path"
	//"strconv"
	"strings"
	//"fmt"
	"io"
)

type Hue struct {
	Host string
	User string
}

type Response struct {
	Success interface{}	`json:"success,omitempty"`
	Error 	*ResponseError			`json:"error,omitempty"`
}

type ResponseError struct {
	Type 		int 	`json:"type,omitempty"`
	Address 	string  `json:"address,omitempty"`
	Description string  `json:"description,omitempty"`
}

func (h *Hue) GetApiUrl(str ...string) string {
	u, err := url.Parse(h.Host)
	if err != nil {
		return ""
	}
	u.Path = path.Join(u.Path, "/api/", h.User)
	for _, p := range str {
		u.Path = path.Join(u.Path, p)
	}
	return u.String()
}


// GET a resource from the url
func (h *Hue) GetResource(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// PUT a resource to a url
func (h *Hue) PutResource(url, data string) (*http.Response, error) {

	body := strings.NewReader(data)

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

	return res, nil

}

// PUT a resource to a url
func (h *Hue) PostResource(url, data string) (*http.Response, error) {

	body := strings.NewReader(data)

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

	return res, nil

}

// DELETE a resource to a url
func (h *Hue) DeleteResource(url string) (*http.Response, error) {

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

	return res, nil

}

// Decode
func (h *Hue) Decode(res io.ReadCloser, in interface{}) error {
	err := json.NewDecoder(res).Decode(&in)
	if err != nil {
		return err
	}
	return nil
}

func New(h, u string) *Hue {
	return &Hue{h, u}
}
