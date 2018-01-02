package huego

import(
  "encoding/json"
  "strconv"
)

type Resourcelink struct {
  Name	string `json:"name,omitempty"`
  Description string `json:"description,omitempty"`
  Type string `json:"type,omitempty"`
  ClassId uint16 `json:"classid,omitempty"`
  Owner string `json:"owner,omitempty"`
  Links []string `json:"links,omitempty"`
  Id int `json:",omitempty"`
}

// Get all resourcelinks
func (h *Hue) GetResourcelinks() ([]*Resourcelink, error) {

  var r map[string]Resourcelink

  res, err := h.GetResource(h.GetApiUrl("/resourcelinks/"))
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &r)
  if err != nil {
    return nil, err
  }

  resourcelinks := make([]*Resourcelink, 0, len(r))

  for i, s := range r {
    s.Id, err = strconv.Atoi(i)
    if err != nil {
      return nil, err
    }
    resourcelinks = append(resourcelinks, &s)
  }

  return resourcelinks, nil

}

// Get one resourcelink
func (h *Hue) GetResourcelink(i int) (*Resourcelink, error) {

	var resourcelink *Resourcelink

	res, err := h.GetResource(h.GetApiUrl("/resourcelinks/", strconv.Itoa(i)))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &resourcelink)
	if err != nil {
		return nil, err
	}

	return resourcelink, nil

}

// Create a resourcelink
func (h *Hue) CreateResourcelink(s *Resourcelink) (*Response, error) {

  var a []*ApiResponse

  data, err := json.Marshal(&s)
  if err != nil {
    return nil, err
  }

  url := h.GetApiUrl("/resourcelinks/")
  res, err := h.PostResource(url, data)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &a)
  if err != nil {
    return nil, err
  }

  resp, err := handleResponse(a)
  if err != nil {
    return nil, err
  }

  return resp, nil

}

// Update a resourcelink
func (h *Hue) UpdateResourcelink(i int, resourcelink *Resourcelink) (*Response, error) {
	var a []*ApiResponse

	data, err := json.Marshal(&resourcelink)
	if err != nil {
		return nil, err
	}

	url := h.GetApiUrl("/resourcelinks/", strconv.Itoa(i))
	res, err := h.PutResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &a)
	if err != nil {
		return nil, err
  }
  
  resp, err := handleResponse(a)
  if err != nil {
    return nil, err
  }

	return resp, nil
}

// Delete a resourcelink
func (h *Hue) DeleteResourcelink(i int) error {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/resourcelinks/", id)

	res, err := h.DeleteResource(url)
	if err != nil {
		return err
	}

	_ = json.Unmarshal(res, &a)
	
  _, err = handleResponse(a)
  if err != nil {
    return err
  }

	return nil
}
