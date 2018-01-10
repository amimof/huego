package huego

import(
  "encoding/json"
  "strconv"
)

// https://developers.meethue.com/documentation/resourcelinks-api
type Resourcelink struct {
  Name	string `json:"name,omitempty"`
  Description string `json:"description,omitempty"`
  Type string `json:"type,omitempty"`
  ClassId uint16 `json:"classid,omitempty"`
  Owner string `json:"owner,omitempty"`
  Links []string `json:"links,omitempty"`
  Id int `json:",omitempty"`
}

// Returns all resourcelinks known to the bridge
func (b *Bridge) GetResourcelinks() ([]*Resourcelink, error) {

  var r map[string]Resourcelink

  url, err := b.getApiPath("/resourcelinks/")
  if err != nil {
    return nil, err
  }

  res, err := b.getResource(url)
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

// Returns one resourcelink by its id defined by i
func (b *Bridge) GetResourcelink(i int) (*Resourcelink, error) {

	var resourcelink *Resourcelink

  url, err := b.getApiPath("/resourcelinks/", strconv.Itoa(i))


	res, err := b.getResource(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &resourcelink)
	if err != nil {
		return nil, err
	}

	return resourcelink, nil

}

// Creates one new resourcelink on the bridge
func (b *Bridge) CreateResourcelink(s *Resourcelink) (*Response, error) {

  var a []*ApiResponse

  data, err := json.Marshal(&s)
  if err != nil {
    return nil, err
  }

  url, err := b.getApiPath("/resourcelinks/")
  if err != nil {
    return nil, err
  }  

  res, err := b.postResource(url, data)
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

// Updates one resourcelink with attributes defined by resourcelink
func (b *Bridge) UpdateResourcelink(i int, resourcelink *Resourcelink) (*Response, error) {
	var a []*ApiResponse

	data, err := json.Marshal(&resourcelink)
	if err != nil {
		return nil, err
	}

  url, err := b.getApiPath("/resourcelinks/", strconv.Itoa(i))
  if err != nil {
    return nil, err
  }

	res, err := b.putResource(url, data)
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

// Deletes one resourcelink with the id of i
func (b *Bridge) DeleteResourcelink(i int) error {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
  url, err := b.getApiPath("/resourcelinks/", id)
  if err != nil {
    return err
  }

	res, err := b.deleteResource(url)
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
