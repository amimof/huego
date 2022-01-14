package huego

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

func TestGetLight(t *testing.T) {

	tests := []struct {
		Body       string
		StatusCode int
		Error      error
		Expect     *Light
	}{
		// Return one light
		{
			Body:       `{"errors":[],"data":[{"alert":{"action_values":["breathe"]},"color":{"gamut":{"blue":{"x":0.1532,"y":0.0475},"green":{"x":0.17,"y":0.7},"red":{"x":0.6915,"y":0.3083}},"gamut_type":"C","xy":{"x":0.4575,"y":0.4099}},"color_temperature":{"mirek":366,"mirek_schema":{"mirek_maximum":500,"mirek_minimum":153},"mirek_valid":true},"dimming":{"brightness":100.0,"min_dim_level":0.20000000298023224},"dynamics":{"speed":0.0,"speed_valid":false,"status":"none","status_values":["none","dynamic_palette"]},"effects":{"effect_values":["no_effect","candle","fire"],"status":"no_effect","status_values":["no_effect","candle","fire"]},"id":"63939343-2449-48b7-aa7d-5daea11dc546","id_v1":"/lights/45","metadata":{"archetype":"spot_bulb","name":"Hue color spot 3"},"mode":"normal","on":{"on":false},"owner":{"rid":"120a0d93-0db0-4b52-8ecb-1ddd64a7e2d9","rtype":"device"},"type":"light"}]}`,
			StatusCode: 200,
			Error:      nil,
			Expect: &Light{
				Alert: &Alert{
					ActionValues: []string{"breathe"},
				},
				Color: &Color{
					Xy: &Xy{
						X: ptrFloat32(0.4575),
						Y: ptrFloat32(0.4099),
					},
					Gamut: &Gamut{
						Blue: &Xy{
							X: ptrFloat32(0.1532),
							Y: ptrFloat32(0.0475),
						},
						Green: &Xy{
							X: ptrFloat32(0.17),
							Y: ptrFloat32(0.7),
						},
						Red: &Xy{
							X: ptrFloat32(0.6915),
							Y: ptrFloat32(0.3083),
						},
					},
					GamutType: ptrString("C"),
				},
				ColorTemperature: &ColorTemperature{
					Mirek: ptrUint16(366),
					MirekSchema: &MirekSchema{
						MirekMaximum: ptrUint16(500),
						MirekMinimum: ptrUint16(153),
					},
					MirekValid: ptrBool(true),
				},
				Dimming: &Dimming{
					Brightness:  ptrFloat32(100),
					MinDimLevel: ptrFloat32(0.20000000298023224),
				},
				Dynamics: &Dynamics{
					Speed:        ptrFloat32(0),
					SpeedValid:   ptrBool(false),
					Status:       ptrString("none"),
					StatusValues: []string{"none", "dynamic_palette"},
				},
				Effects: &Effects{
					EffectValues: []string{"no_effect", "candle", "fire"},
					Status:       ptrString("no_effect"),
					StatusValues: []string{"no_effect", "candle", "fire"},
				},
				Mode: ptrString("normal"),
				On: &On{
					On: ptrBool(false),
				},
				BaseResource: BaseResource{
					Type: ptrString(TypeLight),
					ID:   ptrString("63939343-2449-48b7-aa7d-5daea11dc546"),
					IDv1: ptrString("/lights/45"),
					Metadata: map[string]string{
						"archetype": "spot_bulb",
						"name":      "Hue color spot 3",
					},
					Owner: &Owner{
						Rid:   ptrString("120a0d93-0db0-4b52-8ecb-1ddd64a7e2d9"),
						Rtype: ptrString("device"),
					},
				},
			},
		},
	}

	httpClient := http.DefaultClient
	httpClient.Transport = testTransport
	client, err := NewClient("127.0.0.1", "f0")
	if err != nil {
		t.Fatal(err)
	}
	client.SetClient(httpClient)

	for _, test := range tests {
		testTransport.DoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				Body:       nopCloser(bytes.NewBufferString(test.Body)),
				StatusCode: test.StatusCode,
			}, nil
		}
		light, err := client.GetLight("63939343-2449-48b7-aa7d-5daea11dc546")
		if err != nil {
			t.Fatalf("Unexpected error: %s\n", err.Error())
		}
		if err != nil && err.Error() != test.Error.Error() {
			t.Fatalf("want: %+v\n, got: %+v", test.Error, err)
		}
		if !reflect.DeepEqual(light, test.Expect) {
			t.Fatalf("want: %+v\n, got: %+v", test.Expect, light)
		}
	}

}
