package huego

import(
  "encoding/json"
  "strconv"
)

type Rule struct {
  Name	string `json:"name,omitempty"`
  LastTriggered string `json:"lasttriggered,omitempty"`
  CreationTime string `json:"creationtime,omitempty"`
  TimesTriggered int `json:"timestriggered,omitempty"`
  Owner string `json:"owner,omitempty"`
  Status string `json:"status,omitempty"`
  Conditions []*Condition `json:"conditions,omitempty"`
  Actions []*RuleAction `json:"actions,omitempty"`
  Id int `json:",omitempty"`
}

type Condition struct {
  Address string `json:"address,omitempty"`
  Operator string `json:"operator,omitempty"`
  Value string `json:"string,omitempty"`
}

type RuleAction struct {
  Address string `json:"address,omitempty"`
  Method string `json:"method,omitempty"`
  Body interface{} `json:"body,omitempty"`
}

// Get all rules
func (h *Hue) GetRules() ([]*Rule, error) {

  var r map[string]Rule

  res, err := h.GetResource(h.GetApiUrl("/rules/"))
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &r)
  if err != nil {
    return nil, err
  }

  rules := make([]*Rule, 0, len(r))

  for i, s := range r {
    s.Id, err = strconv.Atoi(i)
    if err != nil {
      return nil, err
    }
    rules = append(rules, &s)
  }

  return rules, nil

}

// Get one rule
func (h *Hue) GetRule(i int) (*Rule, error) {

	var rule *Rule

	res, err := h.GetResource(h.GetApiUrl("/rules/", strconv.Itoa(i)))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &rule)
	if err != nil {
		return nil, err
	}

	return rule, nil

}

// Create a rule
func (h *Hue) CreateRule(s *Rule) (*Response, error) {

  var a []*ApiResponse

  data, err := json.Marshal(&s)
  if err != nil {
    return nil, err
  }

  url := h.GetApiUrl("/rules/")
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

// Update a rule
func (h *Hue) UpdateRule(i int, rule *Rule) (*Response, error) {
  
  var a []*ApiResponse

	data, err := json.Marshal(&rule)
	if err != nil {
		return nil, err
	}

	url := h.GetApiUrl("/rules/", strconv.Itoa(i))
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

// Delete a rule
func (h *Hue) DeleteRule(i int) error {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/rules/", id)

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
