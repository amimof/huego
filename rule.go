package huego

// Rule represents a bridge rule https://developers.meethue.com/documentation/rules-api
type Rule struct {
	Name           string        `json:"name,omitempty"`
	LastTriggered  string        `json:"lasttriggered,omitempty"`
	CreationTime   string        `json:"creationtime,omitempty"`
	TimesTriggered int           `json:"timestriggered,omitempty"`
	Owner          string        `json:"owner,omitempty"`
	Status         string        `json:"status,omitempty"`
	Conditions     []*Condition  `json:"conditions,omitempty"`
	Actions        []*RuleAction `json:"actions,omitempty"`
	ID             int           `json:",omitempty"`
}

// Condition defines the condition of a rule
type Condition struct {
	Address  string `json:"address,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

// RuleAction defines the rule to execute when a rule triggers
type RuleAction struct {
	Address string      `json:"address,omitempty"`
	Method  string      `json:"method,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}
