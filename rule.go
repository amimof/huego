package huego

import(
  "encoding/json"
  "strconv"
)

// https://developers.meethue.com/documentation/rules-api
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

// Returns all rules known to the bridge
func (b *Bridge) GetRules() ([]*Rule, error) {

  var r map[string]Rule

  url, err := b.getApiPath("/rules/")
  if err != nil {
    return nil, err
  }

  res, err := get(url)
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

// Returns one rule by its id of i
func (b *Bridge) GetRule(i int) (*Rule, error) {

	var rule *Rule

  url, err := b.getApiPath("/rules/", strconv.Itoa(i))
  if err != nil {
    return nil, err
  }

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &rule)
	if err != nil {
		return nil, err
	}

	return rule, nil

}

// Creates one rule with attribues defined in s
func (b *Bridge) CreateRule(s *Rule) (*Response, error) {

  var a []*ApiResponse

  data, err := json.Marshal(&s)
  if err != nil {
    return nil, err
  }

  url, err := b.getApiPath("/rules/")
  if err != nil {
    return nil, err
  }
  
  res, err := post(url, data)
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

// Updates one rule by its id of i and rule configuration of rule
func (b *Bridge) UpdateRule(i int, rule *Rule) (*Response, error) {
  
  var a []*ApiResponse

	data, err := json.Marshal(&rule)
	if err != nil {
		return nil, err
	}

  url, err := b.getApiPath("/rules/", strconv.Itoa(i))
  if err != nil {
    return nil, err
  }

	res, err := put(url, data)
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

// Deletes one rule from the bridge
func (b *Bridge) DeleteRule(i int) error {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
  url, err := b.getApiPath("/rules/", id)
  if err != nil {
    return err
  }

	res, err := delete(url)
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
