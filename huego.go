package huego

import (
	//"github.com/amimof/loglevel-go"
	"net/http"
	//"crypto/tls"
	"net/url"
	"path"
	//"fmt"
)

var (
	host string
	client *http.Client
)

type Hue struct {}

type Response struct {
	Success map[string]interface{} 	`json:"success"`
	Error 	*ResponseError			`json:"error"`
}

type ResponseError struct {
	Type 		int 	`json:"type"`
	Address 	string  `json:"address"`
	Description string  `json:description`
}

func GetPath(p string) string {
	basePath := "/api/sca0m37mBcnrEu4jBgMGrdo7uw8rGziaOeMWT9fJ"
	u, err := url.Parse(host)
	if err != nil {
		return ""
	}
	u.Path = path.Join(u.Path, basePath, p)
	return u.String()
}

func New(h string) *Hue {
	host = h
	client = &http.Client{}
	return &Hue{}
}