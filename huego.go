package huego

import (
	//"github.com/amimof/loglevel-go"
	//"net/http"
	//"crypto/tls"

	"net/url"
	"path"
	//"fmt"
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


func New(h, u string) *Hue {
	return &Hue{h, u}
}