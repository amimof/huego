package huego

import (
	"encoding/json"
)

type Room struct {
	Services []*ResourceIdentifier `json:"services,omitempty"`
	Children []*ResourceIdentifier `json:"children,omitempty"`
	BaseResource
}

// Raw marshals the Room into a byte array. Returns nil if errors occur on the way
func (r *Room) Raw() []byte {
	d, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return d
}