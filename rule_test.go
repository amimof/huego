package huego_test

import (
	"testing"

	"github.com/amimof/huego"
	"github.com/stretchr/testify/assert"
)

func TestGetRules(t *testing.T) {
	b := huego.New(hostname, username)
	rules, err := b.GetRules()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d rules", len(rules))
	for _, rule := range rules {
		t.Log(rule)
	}
	assert.Equal(t, "Wall Switch Rule", rules[0].Name)
	assert.Equal(t, "Wall Switch Rule 2", rules[1].Name)
}

func TestGetRule(t *testing.T) {
	b := huego.New(hostname, username)
	l, err := b.GetRule(1)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(l)
	}
}

func TestCreateRule(t *testing.T) {
	b := huego.New(hostname, username)
	conditions := []*huego.Condition{
		{
			Address:  "/sensors/2/state/buttonevent",
			Operator: "eq",
			Value:    "16",
		},
	}
	actions := []*huego.RuleAction{
		{
			Address: "/groups/0/action",
			Method:  "PUT",
			Body:    &huego.State{On: true},
		},
	}
	rule := &huego.Rule{
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
	b := huego.New(hostname, username)
	id := 1
	resp, err := b.UpdateRule(id, &huego.Rule{
		Actions: []*huego.RuleAction{
			{
				Address: "/groups/1/action",
				Method:  "PUT",
				Body:    &huego.State{On: true},
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
	b := huego.New(hostname, username)
	id := 1
	err := b.DeleteRule(id)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Rule %d deleted", id)
	}
}
