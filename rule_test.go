package huego

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRules(t *testing.T) {
	b := New(hostname, username)
	rules, err := b.GetRules()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d rules", len(rules))
	for _, rule := range rules {
		t.Log(rule)
	}

	contains := func(name string, ss []*Rule) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("Wall Switch Rule", rules))
	assert.True(t, contains("Wall Switch Rule 2", rules))
}

func TestGetRule(t *testing.T) {
	b := New(hostname, username)
	l, err := b.GetRule(1)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(l)
	}
}

func TestCreateRule(t *testing.T) {
	b := New(hostname, username)
	conditions := []*Condition{
		{
			Address:  "/sensors/2/state/buttonevent",
			Operator: "eq",
			Value:    "16",
		},
	}
	actions := []*RuleAction{
		{
			Address: "/groups/0/action",
			Method:  "PUT",
			Body:    &State{On: true},
		},
	}
	rule := &Rule{
		Name:       "Huego Test Rule",
		Conditions: conditions,
		Actions:    actions,
	}
	resp, err := b.CreateRule(rule)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Rule created")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestUpdateRule(t *testing.T) {
	b := New(hostname, username)
	id := 1
	resp, err := b.UpdateRule(id, &Rule{
		Actions: []*RuleAction{
			{
				Address: "/groups/1/action",
				Method:  "PUT",
				Body:    &State{On: true},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Rule %d updated", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
}

func TestDeleteRule(t *testing.T) {
	b := New(hostname, username)
	id := 1
	err := b.DeleteRule(id)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Rule %d deleted", id)
	}
}
