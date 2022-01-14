package huego

//type Resource BaseResource

// BaseResource is a type that specifies common attributeees for all resources
type BaseResource struct {
	// +optional
	Type *string `json:"type,omitempty"`
	// +required
	ID *string `json:"id,omitempty"`
	// +optional
	IDv1 *string `json:"id_v1,omitempty"`
	// +required
	Metadata map[string]string `json:"metadata,omitempty"`
	// +required
	Owner *Owner `json:"owner,omitempty"`
}

// Owner is the owner attributes for base resources
type Owner struct {
	// +required
	Rid *string `json:"rid,omitempty"`
	// +required
	Rtype *string `json:"rtype,omitempty"`
}

// Resource is an interface that represents a resource
type Resource interface {
	Type() *string
	Id() *string
	Metadata() map[string]string
	Owner() *Owner
}
