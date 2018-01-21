package huego_test

import (
	"github.com/amimof/huego"
	"os"
	"testing"
)

func TestGetRules(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	rules, err := hue.GetRules()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d rules", len(rules))
	for _, rule := range rules {
		t.Log(rule)
	}
}

func TestGetRule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	rules, err := hue.GetRules()
	if err != nil {
		t.Fatal(err)
	}
	for _, rule := range rules {
		l, err := hue.GetRule(rule.ID)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Log(l)
		}
		break
	}
}

func TestCreateRule(t *testing.T) {
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	conditions := []*huego.Condition{
		&huego.Condition{
			Address:  "/sensors/2/state/buttonevent",
			Operator: "eq",
			Value:    "16",
		},
	}
	actions := []*huego.RuleAction{
		&huego.RuleAction{
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
	resp, err := hue.CreateRule(rule)
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
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	resp, err := hue.UpdateRule(id, &huego.Rule{
		Actions: []*huego.RuleAction{
			&huego.RuleAction{
				Address: "/groups/3/action",
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
	hue := huego.New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	id := 3
	err := hue.DeleteRule(id)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Rule %d deleted", id)
	}
}
