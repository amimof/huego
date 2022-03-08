package huego

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
	"github.com/luci/go-render/render"
)

func TestGetScene(t *testing.T) {

	httpClient := http.DefaultClient
	httpClient.Transport = testTransport
	client, err := NewClientV2("127.0.0.1", "f0")
	if err != nil {
		t.Fatal(err)
	}
	client.SetClient(httpClient)

	tests := []struct {
		Body       string
		StatusCode int
		Error      error
		Expect     interface{}
		Exec			 func() (interface{}, error)
	}{
		// Return one scene
		{
			Body:       `{"errors":[],"data":[{"actions":[{"action":{"color_temperature":{"mirek":387},"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"66618474-95d7-4413-92e7-4eae71a80f1c","rtype":"light"}},{"action":{"color":{"xy":{"x":0.5897,"y":0.3536}},"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"466af769-8a20-420d-ad79-f92ffb1e72e5","rtype":"light"}},{"action":{"gradient":{"points":[{"xy":{"x":0.5897,"y":0.3536}},{"xy":{"x":0.6563,"y":0.3211}}]},"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"fbe3d0b3-bf5c-4c79-a429-84571c1dc8fc","rtype":"light"}}],"group":{"rid":"17853064-4cd0-431e-8f7c-f19640b046fb","rtype":"zone"},"id":"2f9ca1a6-7b96-4ac3-98cd-7a4c5be5bb30","id_v1":"/scenes/lXY9hGKn1-9dxlO","metadata":{"image":{"rid":"4f2ed241-5aea-4c9d-8028-55d2b111e06f","rtype":"public_image"},"name":"Savannasunset"},"palette":{"color":[{"color":{"xy":{"x":0.6563,"y":0.3211}},"dimming":{"brightness":80.71}}],"color_temperature":[{"color_temperature":{"mirek":373},"dimming":{"brightness":80.71}}],"dimming":[{"brightness":80.71}]},"speed":0.6190476190476191,"type":"scene"}]}`,
			StatusCode: 200,
			Error:      nil,
			Exec: func() (interface{}, error) {
				return client.GetScene("2f9ca1a6-7b96-4ac3-98cd-7a4c5be5bb30")
			},
			Expect: &Scene{
				Actions: []*ActionGet{
					{
						Action: &Action{
							ColorTemperature: NewColorTemperature(387),
							Dimming:          NewDimming(77.95),
							On: &On{
								On: ptrBool(true),
							},
						},
						Target: &Owner{
							Rid:   ptrString("66618474-95d7-4413-92e7-4eae71a80f1c"),
							Rtype: ptrString("light"),
						},
					},
					{
						Action: &Action{
							Color: NewColor(0.5897, 0.3536),
							Dimming: &Dimming{
								Brightness: ptrFloat32(77.95),
							},
							On: &On{
								On: ptrBool(true),
							},
						},
						Target: &Owner{
							Rid:   ptrString("466af769-8a20-420d-ad79-f92ffb1e72e5"),
							Rtype: ptrString("light"),
						},
					},
					{
						Action: &Action{
							Gradient: NewGradient(NewColor(0.5897, 0.3536), NewColor(0.6563, 0.3211)),
							Dimming: &Dimming{
								Brightness: ptrFloat32(77.95),
							},
							On: &On{
								On: ptrBool(true),
							},
						},
						Target: &Owner{
							Rid:   ptrString("fbe3d0b3-bf5c-4c79-a429-84571c1dc8fc"),
							Rtype: ptrString("light"),
						},
					},
				},
				Group: &Owner{
					Rid:   ptrString("17853064-4cd0-431e-8f7c-f19640b046fb"),
					Rtype: ptrString("zone"),
				},
				BaseResource: BaseResource{
					Type: ptrString("scene"),
					ID:   ptrString("2f9ca1a6-7b96-4ac3-98cd-7a4c5be5bb30"),
					IDv1: ptrString("/scenes/lXY9hGKn1-9dxlO"),
					Metadata: map[string]interface{}{
						"image": map[string]interface{}{
							"rid":   "4f2ed241-5aea-4c9d-8028-55d2b111e06f",
							"rtype": "public_image",
						},
						"name": "Savannasunset",
					},
				},
				Palette: &Palette{
					Color: []*PaletteColor{
						{
							Color:   NewColor(0.6563, 0.3211),
							Dimming: NewDimming(80.71),
						},
					},
					ColorTemperature: []*PaletteColorTemperature{
						{
							ColorTemperature: NewColorTemperature(373),
							Dimming:          NewDimming(80.71),
						},
					},
					Dimming: []*Dimming{NewDimming(80.71)},
				},
				Speed: ptrFloat64(0.6190476190476191),
			},
		},
	}

	for _, test := range tests {
		testTransport.DoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				Body:       nopCloser(bytes.NewBufferString(test.Body)),
				StatusCode: test.StatusCode,
			}, nil
		}
		obj, err := test.Exec()
		if err != nil {
			t.Fatalf("Unexpected error: %s\n", err.Error())
		}
		if err != nil && err.Error() != test.Error.Error() {
			t.Fatalf("want: %+v\n, got: %+v", test.Error, err)
		}
		if !reflect.DeepEqual(obj, test.Expect) {
			t.Fatalf("want: %+v\n, got: %s", render.Render(test.Expect), render.Render(obj))
		}
	}

}
