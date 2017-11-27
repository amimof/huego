package huego

import (
	"testing"
	"os"
)

func TestGetRules(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	rules, err := hue.GetRules()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Found %d rules", len(rules))
	for _, rule := range rules {
		t.Log(rule)
	}
}


func TestGetRule(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	rules, err := hue.GetRules()
	if err != nil {
		t.Error(err)
	}
	for _, rule := range rules {
		l, err := hue.GetRule(rule.Id)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(l)
		}
		break
	}
}

func TestCreateRule(t *testing.T) {
	hue := New(os.Getenv("HUE_HOSTNAME"), os.Getenv("HUE_USERNAME"))
	condition := []&Condition{{
		Address: "/sensors/2/state/buttonevent",
		Operator: "eq"
		Value: "16",
	}}
	action := []&Action{{
		Address: "/groups/0/action",
		Method: "PUT",
		Body:
	}}
	response, err := hue.CreateRule(rule)
	if err != nil {
		t.Error(err)
	}
	for _, r := range response {
		t.Logf("Response from put: Success=%v Error=%v", r.Success, r.Error)
	}
}

// func TestDeleteRule(t *testing.T) {
// 	res, err := hue.DeleteRule(1)
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	} else {
// 		for _, r := range res {
// 			t.Log(r.Success, r.Error)
// 		}
// 	}
// }
