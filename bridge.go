package huego

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

type Bridge struct {
	Host string `json:"internalipaddress,omitempty"` 
	User string
	Id string `json:"id,omitempty"`
}

func (b *Bridge) getApiPath(str ...string) (string, error) {

	if strings.Index(strings.ToLower(b.Host), "http://") <= -1 && strings.Index(strings.ToLower(b.Host), "https://") <= -1 {
		b.Host = fmt.Sprintf("%s%s", "http://", b.Host)
	} 

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

// Calls New() and passes Host on this Bridge instance
func (b *Bridge) Login(u string) *Bridge {
	return New(b.Host, u)
}