package huego

import (
	"encoding/json"
)

type Zone struct {
	Services []*ResourceIdentifier `json:"services,omitempty"`
	Items []*ResourceIdentifier `json:"items,omitempty"`
	Children []*ResourceIdentifier `json:"children,omitempty"`
}

// Raw marshals the Room into a byte array. Returns nil if errors occur on the way
func (z *Zone) Raw() []byte {
	d, err := json.Marshal(z)
	if err != nil {
		return nil
	}
	return d
}