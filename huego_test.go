package huego

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// I'm too lazy to have this elsewhere
// export HUE_USERNAME=9D6iHMbM-Bt7Kd0Cwh9Quo4tE02FMnmrNrnFAdAq
// export HUE_HOSTNAME=192.168.1.59

var username string
var hostname string
var badHostname = "bad-hue-config"

func init() {

	hostname = os.Getenv("HUE_HOSTNAME")
	username = os.Getenv("HUE_USERNAME")

	tests := []struct {
		method string
		path   string
		data   string
		url    string
	}{
		// DISCOVERY
		{
			method: "GET",
			url:    "https://discovery.meethue.com",
			data:   `[{"id":"001788fffe73ff19","internalipaddress":"192.168.13.112"}]`,
		},

		// CONFIG
		{
			method: "GET",
			path:   "/config",
			data:   `{"name":"Philipshue","zigbeechannel":15,"mac":"00:17:88:00:00:00","dhcp":true,"ipaddress":"192.168.1.7","netmask":"255.255.255.0","gateway":"192.168.1.1","proxyaddress":"none","proxyport":0,"UTC":"2014-07-17T09:27:35","localtime":"2014-07-17T11:27:35","timezone":"Europe/Madrid","whitelist":{"ffffffffe0341b1b376a2389376a2389":{"lastusedate":"2014-07-17T07:21:38","createdate":"2014-04-08T08:55:10","name":"PhilipsHueAndroidApp#TCTALCATELONETOU"},"pAtwdCV8NZId25Gk":{"lastusedate":"2014-05-07T18:28:29","createdate":"2014-04-09T17:29:16","name":"MyApplication"},"gDN3IaPYSYNPWa2H":{"lastusedate":"2014-05-07T09:15:21","createdate":"2014-05-07T09:14:38","name":"iPhoneWeb1"}},"swversion":"01012917","apiversion":"1.3.0","swupdate":{"updatestate":2,"checkforupdate":false,"devicetypes":{"bridge":true,"lights":["1","2","3"], "sensors":["4","5","6"]},"url":"","text":"010000000","notify":false},"linkbutton":false,"portalservices":false,"portalconnection":"connected","portalstate":{"signedon":true,"incoming":false,"outgoing":true,"communication":"disconnected"}}`,
		},
		{
			method: "PUT",
			path:   "/config",
			data:   `[{"success":{"/config/name":"My bridge"}}]`,
		},
		{
			method: "GET",
			path:   path.Join("/invalid_password", "/lights"),
			data:   `[{"error":{"type":1,"address":"/lights","description":"unauthorized user"}}]`,
		},
		{
			method: "GET",
			path:   "",
			data:   `{"lights":{"1":{"state":{"on":false,"bri":0,"hue":0,"sat":0,"xy":[0.0000,0.0000],"ct":0,"alert":"none","effect":"none","colormode":"hs","reachable":true},"type":"Extendedcolorlight","name":"HueLamp1","modelid":"LCT001","swversion":"65003148"},"2":{"state":{"on":true,"bri":254,"hue":33536,"sat":144,"xy":[0.3460,0.3568],"ct":201,"alert":"none","effect":"none","colormode":"hs","reachable":true},"type":"Extendedcolorlight","name":"HueLamp2","modelid":"LCT001","swversion":"65003148"}},"groups":{"1":{"action":{"on":true,"bri":254,"hue":33536,"sat":144,"xy":[0.3460,0.3568],"ct":201,"effect":"none","colormode":"xy"},"lights":["1","2"],"name":"Group1"}},"config":{"name":"Philipshue","mac":"00:00:88:00:bb:ee","dhcp":true,"ipaddress":"192.168.1.74","netmask":"255.255.255.0","gateway":"192.168.1.254","proxyaddress":"","proxyport":0,"UTC":"2012-10-29T12:00:00","whitelist":{"1028d66426293e821ecfd9ef1a0731df":{"lastusedate":"2012-10-29T12:00:00","createdate":"2012-10-29T12:00:00","name":"testuser"}},"swversion":"01003372","swupdate":{"updatestate":2,"checkforupdate":false,"devicetypes":{"bridge":true,"lights":["1","2","3"], "sensors":["4","5","6"]},"url":"","text":"010000000","notify":false},"swupdate2":{"checkforupdate":false,"lastchange":"2017-06-21T19:44:36","bridge":{"state":"noupdates","lastinstall":"2017-06-21T19:44:18"},"state":"noupdates","autoinstall":{"updatetime":"T14:00:00","on":false}},"schedules":{"1":{"name":"schedule","description":"","command":{"address":"/api/<username>/groups/0/action","body":{"on":true},"method":"PUT"},"time":"2012-10-29T12:00:00"}}}}`,
		},
		{
			method: "POST",
			path:   "",
			data:   `[{"success":{"username": "83b7780291a6ceffbe0bd049104df", "clientkey": "33DDF493992908E3D97AAAA5A6F189E1"}}]`,
		},
		{
			method: "DELETE",
			path:   "/config/whitelist/ffffffffe0341b1b376a2389376a2389",
			data:   `[{"success": "/config/whitelist/ffffffffe0341b1b376a2389376a2389 deleted."}]`,
		},

		// LIGHT
		{
			method: "GET",
			path:   path.Join(username, "/lights"),
			data:   `{"1":{"state":{"on":false,"bri":1,"hue":33761,"sat":254,"effect":"none","xy":[0.3171,0.3366],"ct":159,"alert":"none","colormode":"xy","mode":"homeautomation","reachable":true},"swupdate":{"state":"noupdates","lastinstall":"2017-10-15T12:07:34"},"type":"Extendedcolorlight","name":"Huecolorlamp7","modelid":"LCT001","manufacturername":"Philips","productname":"Huecolorlamp","capabilities":{"certified":true,"control":{"mindimlevel":5000,"maxlumen":600,"colorgamuttype":"B","colorgamut":[[0.675,0.322],[0.409,0.518],[0.167,0.04]],"ct":{"min":153,"max":500}},"streaming":{"renderer":true,"proxy":false}},"config":{"archetype":"sultanbulb","function":"mixed","direction":"omnidirectional"},"uniqueid":"00:17:88:01:00:bd:c7:b9-0b","swversion":"5.105.0.21169"},"2":{"state":{"on":false,"bri":1,"hue":35610,"sat":237,"effect":"none","xy":[0.1768,0.395],"ct":153,"alert":"none","colormode":"xy","mode":"homeautomation","reachable":true},"swupdate":{"state":"noupdates","lastinstall":"2017-10-18T12:50:40"},"type":"Extendedcolorlight","name":"Huelightstripplus1","modelid":"LST002","manufacturername":"Philips","productname":"Huelightstripplus","capabilities":{"certified":true,"control":{"mindimlevel":40,"maxlumen":1600,"colorgamuttype":"C","colorgamut":[[0.6915,0.3083],[0.17,0.7],[0.1532,0.0475]],"ct":{"min":153,"max":500}},"streaming":{"renderer":true,"proxy":true}},"config":{"archetype":"huelightstrip","function":"mixed","direction":"omnidirectional"},"uniqueid":"00:17:88:01:02:15:97:46-0b","swversion":"5.105.0.21169"}}`,
		},
		{
			method: "GET",
			path:   path.Join(username, "/lights/1"),
			data:   `{"state":{"on":false,"bri":1,"hue":12594,"sat":251,"effect":"none","xy":[0.5474,0.4368],"alert":"none","colormode":"xy","mode":"homeautomation","reachable":true},"swupdate":{"state":"noupdates","lastinstall":"2017-12-13T01:59:13"},"type":"Colorlight","name":"Huebloom1","modelid":"LLC011","manufacturername":"Philips","productname":"Huebloom","capabilities":{"certified":true,"control":{"mindimlevel":10000,"maxlumen":120,"colorgamuttype":"A","colorgamut":[[0.704,0.296],[0.2151,0.7106],[0.138,0.08]]},"streaming":{"renderer":true,"proxy":false}},"config":{"archetype":"huebloom","function":"decorative","direction":"upwards"},"uniqueid":"00:17:88:01:00:c5:3b:e3-0b","swversion":"5.105.1.21778"}`,
		},
		{
			method: "GET",
			path:   path.Join(username, "/lights/new"),
			data:   `{"7":{"name":"HueLamp7"},"8":{"name":"HueLamp8"},"lastscan":"2012-10-29T12:00:00"}`,
		},
		{
			method: "POST",
			path:   path.Join(username, "/lights"),
			data:   `[{"success":{"/lights":"Searching for new devices"}}]`,
		},
		{
			method: "PUT",
			path:   path.Join(username, "/lights/1/state"),
			data:   `[{"success":{"/lights/1/state/bri":200}},{"success":{"/lights/1/state/on":true}},{"success":{"/lights/1/state/hue":50000}}]`,
		},
		{
			method: "PUT",
			path:   path.Join(username, "/lights/1"),
			data:   `[{"success":{"/lights/1/name":"Bedroom Light"}}]`,
		},
		{
			method: "DELETE",
			path:   path.Join(username, "/lights/1"),
			data:   `[{"success":"/lights/<id> deleted"}]`,
		},

		// GROUP
		{
			method: "GET",
			path:   "/groups",
			data:   `{"1":{"name":"Group 1","lights":["1","2"],"type":"LightGroup","state":{"all_on":true,"any_on":true},"action":{"on":true,"bri":254,"hue":10000,"sat":254,"effect":"none","xy":[0.5,0.5],"ct":250,"alert":"select","colormode":"ct"}},"2":{"name":"Group 2","lights":["3","4","5"],"type":"LightGroup","state":{"all_on":true,"any_on":true},"action":{"on":true,"bri":153,"hue":4345,"sat":254,"effect":"none","xy":[0.5,0.5],"ct":250,"alert":"select","colormode":"ct"}}}`,
		},
		{
			method: "GET",
			path:   "/groups/1",
			data:   `{"action":{"on":true,"hue":0,"effect":"none","bri":100,"sat":100,"ct":500,"xy":[0.5,0.5]},"lights":["1","2"],"state":{"any_on":true,"all_on":true},"type":"Room","class":"Bedroom","name":"Masterbedroom"}`,
		},
		{
			method: "PUT",
			path:   "/groups/1",
			data:   `[{"success":{"/groups/1/lights":["1"]}},{"success":{"/groups/1/name":"Bedroom"}}]`,
		},
		{
			method: "PUT",
			path:   "/groups/1/action",
			data:   `[{"success":{"address":"/groups/1/action/on","value":true}},{"success":{"address":"/groups/1/action/effect","value":"colorloop"}},{"success":{"address":"/groups/1/action/hue","value":6000}}]`,
		},
		{
			method: "POST",
			path:   "/groups",
			data:   `[{"success":{"id":"1"}}]`,
		},
		{
			method: "DELETE",
			path:   "/groups/1",
			data:   `[{"success":"/groups/1 deleted."}]`,
		},

		// SCENE
		{
			method: "GET",
			path:   "/scenes",
			data:   `{"4e1c6b20e-on-0":{"name":"Kathyon1449133269486","lights":["2","3"],"owner":"ffffffffe0341b1b376a2389376a2389","recycle":true,"locked":false,"appdata":{},"picture":"","lastupdated":"2015-12-03T08:57:13","version":1},"3T2SvsxvwteNNys":{"name":"Cozydinner","type":"GroupScene","group":"1","lights":["1","2"],"owner":"ffffffffe0341b1b376a2389376a2389","recycle":true,"locked":false,"appdata":{"version":1,"data":"myAppData"},"picture":"","lastupdated":"2015-12-03T10:09:22","version":2}}`,
		},
		{
			method: "GET",
			path:   "/scenes/4e1c6b20e-on-0",
			data:   `{"name":"Cozy dinner","type":"GroupScene","group":"1","lights":["1"],"owner":"newdeveloper","recycle":true,"locked":false,"appdata":{},"picture":"","lastupdated":"2015-12-03T10:09:22","version":2,"lightstates":{"1":{"on":true,"bri":237,"xy":[0.5806,0.3903]}}}`,
		},
		{
			method: "POST",
			path:   "/scenes",
			data:   `[{"success":{"address":"/scenes/ab341ef24/name","value":"Romanticdinner"}},{"success":{"address":"/scenes/ab3C41ef24/lights","value":["1","2"]}}]`,
		},
		{
			method: "PUT",
			path:   "/scenes/4e1c6b20e-on-0",
			data:   `[{"success":{"/scenes/74bc26d5f-on-0/name":"Cozydinner"}},{"success":{"/scenes/74bc26d5f-on-0/storelightstate":true}},{"success":{"/scenes/74bc26d5f-on-0/lights":["2","3"]}}]`,
		},
		{
			method: "PUT",
			path:   "/scenes/4e1c6b20e-on-0/lightstates/1",
			data:   `[{"success":{"address":"/scenes/ab341ef24/lights/1/state/on","value":true}},{"success":{"address":"/scenes/ab341ef24/lights/1/state/ct","value":200}}]`,
		},
		{
			method: "DELETE",
			path:   "/scenes/4e1c6b20e-on-0",
			data:   `[{"success":"/scenes/3T2SvsxvwteNNys deleted"}]`,
		},

		// RULE
		{
			method: "GET",
			path:   "/rules",
			data:   `{ "1": { "name": "Wall Switch Rule", "lasttriggered": "2013-10-17T01:23:20", "creationtime": "2013-10-10T21:11:45", "timestriggered": 27, "owner": "78H56B12BA", "status": "enabled", "conditions": [ { "address": "/sensors/2/state/buttonevent", "operator": "eq", "value": "16" }, { "address": "/sensors/2/state/lastupdated", "operator": "dx" } ], "actions": [ { "address": "/groups/0/action", "method": "PUT", "body": { "scene": "S3" } } ] }, "2": { "name": "Wall Switch Rule 2" }} `,
		},
		{
			method: "GET",
			path:   "/rules/1",
			data:   `{ "name": "Wall Switch Rule", "owner": "ruleOwner", "created": "2014-07-23T15:02:56", "lasttriggered": "none", "timestriggered": 0, "status": "enabled", "conditions": [ { "address": "/sensors/2/state/buttonevent", "operator": "eq", "value": "16" }, { "address": "/sensors/2/state/lastupdated", "operator": "dx" } ], "actions": [ { "address": "/groups/0/action", "method": "PUT", "body": { "scene": "S3" } } ] }`,
		},
		{
			method: "POST",
			path:   "/rules",
			data:   `[{"success":{"id": "3"}}]`,
		},
		{
			method: "PUT",
			path:   "/rules/1",
			data:   `[ { "success": { "/rules/1/actions": [ { "address": "/groups/0/action", "method": "PUT", "body": { "scene": "S3" } } ] } } ]`,
		},
		{
			method: "DELETE",
			path:   "/rules/1",
			data:   `[{"success": "/rules/1 deleted."}]`,
		},

		// SCHEDULE
		{
			method: "GET",
			path:   "/schedules",
			data:   `{ "1": { "name": "Timer", "description": "", "command": { "address": "/api/s95jtYH8HUVWNkCO/groups/0/action", "body": { "scene": "02b12e930-off-0" }, "method": "PUT" }, "time": "PT00:01:00", "created": "2014-06-23T13:39:16", "status": "disabled", "autodelete": false, "starttime": "2014-06-23T13:39:16" }, "2": { "name": "Alarm", "description": "", "command": { "address": "/api/s95jtYH8HUVWNkCO/groups/0/action", "body": { "scene": "02b12e930-off-0" }, "method": "PUT" }, "localtime": "2014-06-23T19:52:00", "time": "2014-06-23T13:52:00", "created": "2014-06-23T13:38:57", "status": "disabled", "autodelete": false } }`,
		},
		{
			method: "GET",
			path:   "/schedules/1",
			data:   `{ "name": "Wake up", "description": "My wake up alarm", "command": { "address": "/api/<username>/groups/1/action", "method": "PUT", "body": { "on": true } }, "time": "W124/T06:00:00" }`,
		},
		{
			method: "POST",
			path:   "/schedules",
			data:   `[{"success":{"id": "2"}}]`,
		},
		{
			method: "PUT",
			path:   "/schedules/1",
			data:   `[{ "success": {"/schedules/1/name": "Wake up"}}]`,
		},
		{
			method: "DELETE",
			path:   "/schedules/1",
			data:   `[{"success": "/schedules/1 deleted."}]`,
		},

		// SENSOR
		{
			method: "GET",
			path:   "/sensors",
			data:   `{ "1": { "state": { "daylight": false, "lastupdated": "2014-06-27T07:38:51" }, "config": { "on": true, "long": "none", "lat": "none", "sunriseoffset": 50, "sunsetoffset": 50 }, "name": "Daylight", "type": "Daylight", "modelid": "PHDL00", "manufacturername": "Philips", "swversion": "1.0" }, "2": { "state": { "buttonevent": 0, "lastupdated": "none" }, "config": { "on": true }, "name": "Tap Switch 2", "type": "ZGPSwitch", "modelid": "ZGPSWITCH", "manufacturername": "Philips", "uniqueid": "00:00:00:00:00:40:03:50-f2" } }`,
		},
		{
			method: "GET",
			path:   "/sensors/1",
			data:   `{ "state":{ "buttonevent": 34, "lastupdated":"2013-03-25T13:32:34" }, "name": "Wall tap 1", "modelid":"ZGPSWITCH", "uniqueid":"01:23:45:67:89:AB-12", "manufacturername": "Philips", "swversion":"1.0", "type":  "ZGPSwitch" }`,
		},
		{
			method: "POST",
			path:   "/sensors",
			data:   `[ { "success": { "/sensors": "Searching for new devices"}}]`,
		},
		{
			method: "POST",
			path:   "/sensors",
			data:   `[{"success":{"id": "4"}}]`,
		},
		{
			method: "GET",
			path:   "/sensors/new",
			data:   `{ "7": {"name": "Hue Tap 1"}, "8": {"name": "Button 3"}, "lastscan":"2013-05-22T10:24:00" }`,
		},
		{
			method: "PUT",
			path:   "/sensors/1",
			data:   `[{"success":{"/sensors/2/name":"Bedroom Tap"}}]`,
		},
		{
			method: "DELETE",
			path:   "/sensors/1",
			data:   `[{"success": "/sensors/1 deleted."}]`,
		},
		{
			method: "PUT",
			path:   "/sensors/1/config",
			data:   `[{"success":{"/sensors/2/config/on":true}}]`,
		},
		{
			method: "PUT",
			path:   "/sensors/1/state",
			data:   `[{"success":{"/sensors/1/state/presence": false}}]`,
		},

		// CAPABILITIES
		{
			method: "GET",
			path:   "/capabilities",
			data:   `{ "lights":{ "available": 10 }, "sensors":{ "available": 60, "clip": { "available": 60 }, "zll": { "available": 60 }, "zgp": { "available": 60 } }, "groups": {}, "scenes": { "available": 100, "lightstates": { "available": 1500 } }, "rules": {}, "schedules": {}, "resourcelinks": {}, "whitelists": {}, "timezones": { "values":[ "Africa/Abidjan", "Africa/Accra", "Pacific/Wallis", "US/Pacific-New" ] } }`,
		},

		// RESOURCELINK
		{
			method: "GET",
			path:   "/resourcelinks",
			data:   `{ "1": { "name": "Sunrise", "description": "Carla's wakeup experience", "class": 1, "owner": "78H56B12BAABCDEF", "links": ["/schedules/2", "/schedules/3", "/scenes/ABCD", "/scenes/EFGH", "/groups/8"] }, "2": { "name": "Sunrise 2" } }`,
		},
		{
			method: "GET",
			path:   "/resourcelinks/1",
			data:   `{ "name": "Sunrise", "description": "Carla's wakeup experience", "type":"Link", "class": 1, "owner": "78H56B12BAABCDEF", "links": ["/schedules/2", "/schedules/3", "/scenes/ABCD", "/scences/EFGH", "/groups/8"] }`,
		},
		{
			method: "POST",
			path:   "/resourcelinks",
			data:   `[{"success":{"id": "3"}}]`,
		},
		{
			method: "PUT",
			path:   "/resourcelinks/1",
			data:   `[{ "success": { "/resourcelinks/1/name": "Sunrise" } }, { "success": { "/resourcelinks/1/description": "Carla's wakeup experience" } }]`,
		},
		{
			method: "DELETE",
			path:   "/resourcelinks/1",
			data:   `[{"success": "/resourcelinks/1 deleted."}]`,
		},
	}

	httpmock.Activate()

	for _, test := range tests {
		if test.url == "" {
			test.url = fmt.Sprintf("http://%s/api%s", hostname, test.path)
		}
		httpmock.RegisterResponder(test.method, test.url, httpmock.NewStringResponder(200, test.data))
	}

	// Register a responder for bad requests
	paths := []string{
		"",
		"config",
		"/config",
		"/config",
		"/config/whitelist/ffffffffe0341b1b376a2389376a2389",
		path.Join("/invalid_password", "/lights"),
		path.Join(username, "/lights"),
		path.Join(username, "/lights/1"),
		path.Join(username, "/lights/new"),
		path.Join(username, "/lights"),
		path.Join(username, "/lights/1/state"),
		path.Join(username, "/lights/1"),
		path.Join(username, "/lights/1"),
		"/groups",
		"/groups/1",
		"/groups/1",
		"/groups/1/action",
		"/groups",
		"/groups/1",
		"/scenes",
		"/scenes/4e1c6b20e-on-0",
		"/scenes",
		"/scenes/4e1c6b20e-on-0",
		"/scenes/4e1c6b20e-on-0/lightstates/1",
		"/scenes/4e1c6b20e-on-0",
		"/rules",
		"/rules/1",
		"/rules",
		"/rules/1",
		"/rules/1",
		"/schedules",
		"/schedules/1",
		"/schedules",
		"/schedules/1",
		"/schedules/1",
		"/sensors",
		"/sensors/1",
		"/sensors",
		"/sensors",
		"/sensors/new",
		"/sensors/1",
		"/sensors/1",
		"/sensors/1/config",
		"/sensors/1/state",
		"/capabilities",
		"/resourcelinks",
		"/resourcelinks/1",
		"/resourcelinks",
		"/resourcelinks/1",
		"/resourcelinks/1",
	}

	// Register responder for errors
	for _, p := range paths {
		response := []byte("not json")
		httpmock.RegisterResponder("GET", fmt.Sprintf("http://%s/api%s", badHostname, p), httpmock.NewBytesResponder(200, response))
		httpmock.RegisterResponder("POST", fmt.Sprintf("http://%s/api%s", badHostname, p), httpmock.NewBytesResponder(200, response))
		httpmock.RegisterResponder("PUT", fmt.Sprintf("http://%s/api%s", badHostname, p), httpmock.NewBytesResponder(200, response))
		httpmock.RegisterResponder("DELETE", fmt.Sprintf("http://%s/api%s", badHostname, p), httpmock.NewBytesResponder(200, response))
	}

}

func TestDiscoverAndLogin(t *testing.T) {
	bridge, err := Discover()
	if err != nil {
		t.Fatal(err)
	}
	bridge = bridge.Login(username)
	t.Logf("Successfully logged in to bridge")
}

func TestDiscoverAllBridges(t *testing.T) {
	bridges, err := DiscoverAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Discovered %d bridge(s)", len(bridges))
	for i, bridge := range bridges {
		t.Logf("%d: ", i)
		t.Logf("  Host: %s", bridge.Host)
		t.Logf("  User: %s", bridge.User)
		t.Logf("  ID: %s", bridge.ID)
	}
}

func Test_unmarshalError(t *testing.T) {
	s := struct {
		Name string `json:"name"`
	}{
		Name: "amimof",
	}
	err := unmarshal([]byte(`not json`), s)
	assert.NotNil(t, err)
}
