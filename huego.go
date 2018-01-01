package huego

import (
	"net/http"
	"encoding/json"
	"net/url"
	"path"
	"strings"
	"io"
	"io/ioutil"
	"fmt"
)

type Hue struct {
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

// func handleResponse(a []*ApiResponse) (*Response, error) {
// 	var resp []*Response
// 	for _, r := range a {	
// 		if r.Error != nil {
// 			return nil, r.Error
// 		}
// 		if r.Success != nil {
// 			for k, _ := range r.Success {
// 				j, _ := json.Marshal(&r.Success)
// 				resp = append(resp, &Response{
// 					Address: k,
// 					Value: r.Success[k],
// 					Interface: r.Success[k],
// 					Json: j,
// 				})
// 			}
// 		}
// 	}
// 	return resp, nil
// }

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
func (h *Hue) GetResource(url string) ([]byte, error) {

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

// PUT a resource to a url
func (h *Hue) PutResource(url string, data []byte) ([]byte, error) {

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

// PUT a resource to a url
func (h *Hue) PostResource(url string, data []byte) ([]byte, error) {

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

// DELETE a resource to a url
func (h *Hue) DeleteResource(url string) ([]byte, error) {

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
