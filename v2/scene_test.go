package huego

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

func TestGetScene(t *testing.T) {

	tests := []struct {
		Body       string
		StatusCode int
		Error      error
		Expect     *Scene
	}{
		// Return one scene
		{
			Body:       `{"errors":[],"data":[{"actions":[{"action":{"color_temperature":{"mirek":387},"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"66618474-95d7-4413-92e7-4eae71a80f1c","rtype":"light"}},{"action":{"color_temperature":{"mirek":387},"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"b368ea65-482e-4466-b717-266baac2a616","rtype":"light"}},{"action":{"color_temperature":{"mirek":387},"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"131ff5cb-2d46-4837-9a0e-bc59cea44170","rtype":"light"}},{"action":{"color_temperature":{"mirek":387},"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"1298a8c0-0775-4107-96e7-a5cae27aeef5","rtype":"light"}},{"action":{"on":{"on":true}},"target":{"rid":"31dbd6a4-5555-4dff-9b4c-ab5b7434f903","rtype":"light"}},{"action":{"color":{"xy":{"x":0.5897,"y":0.3536}},"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"466af769-8a20-420d-ad79-f92ffb1e72e5","rtype":"light"}},{"action":{"color":{"xy":{"x":0.5897,"y":0.3536}},"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"fbe3d0b3-bf5c-4c79-a429-84571c1dc8fc","rtype":"light"}},{"action":{"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"c63f05d9-2a1f-4918-b602-82f798954599","rtype":"light"}},{"action":{"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"d61d13a9-b252-4209-96b7-a27b31c41bfb","rtype":"light"}},{"action":{"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"54bf9070-cf34-4556-a010-298b35c969bd","rtype":"light"}},{"action":{"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"2e65cf17-6afc-4a84-9de1-9527319b9fb7","rtype":"light"}},{"action":{"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"01544696-dd85-4ff7-b250-626a3599dd9a","rtype":"light"}},{"action":{"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"acce9b3e-b02c-463f-9e63-45f5907b1dfb","rtype":"light"}},{"action":{"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"4d080a60-1b02-40ad-bd8c-a542c8299c69","rtype":"light"}},{"action":{"dimming":{"brightness":77.95},"on":{"on":true}},"target":{"rid":"3dc5a20f-c80b-4abc-a933-876b70f746cd","rtype":"light"}}],"group":{"rid":"17853064-4cd0-431e-8f7c-f19640b046fb","rtype":"zone"},"id":"2f9ca1a6-7b96-4ac3-98cd-7a4c5be5bb30","id_v1":"/scenes/lXY9hGKn1-9dxlO","metadata":{"image":{"rid":"4f2ed241-5aea-4c9d-8028-55d2b111e06f","rtype":"public_image"},"name":"Savanna sunset"},"palette":{"color":[{"color":{"xy":{"x":0.6563,"y":0.3211}},"dimming":{"brightness":80.71}},{"color":{"xy":{"x":0.5862,"y":0.3575}},"dimming":{"brightness":80.71}},{"color":{"xy":{"x":0.5502,"y":0.3655}},"dimming":{"brightness":80.71}},{"color":{"xy":{"x":0.4577,"y":0.4563}},"dimming":{"brightness":80.71}},{"color":{"xy":{"x":0.4162,"y":0.4341}},"dimming":{"brightness":80.71}}],"color_temperature":[{"color_temperature":{"mirek":373},"dimming":{"brightness":80.71}}],"dimming":[{"brightness":80.71}]},"speed":0.6190476190476191,"type":"scene"}]}`,
			StatusCode: 200,
			Error:      nil,
			Expect: &Scene{
				Actions: []*ActionGet{
					&ActionGet{
						Action: &Action{
							ColorTemperature: &ColorTemperature{
								Mirek: ptrUint16(387),
							},
							Dimming: &Dimming{
								Brightness: ptrFloat32(77.95),
							},
							On: &On{
								On: ptrBool(true),
							},
						},
					},
				},
				Group: &Owner{
					Rid: ptrString("66618474-95d7-4413-92e7-4eae71a80f1c"),
					Rtype: ptrString("light"),
				},
			},
		},
	}

	httpClient := http.DefaultClient
	httpClient.Transport = testTransport
	client, err := NewClientV2("127.0.0.1", "f0")
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
		light, err := client.GetScene("2f9ca1a6-7b96-4ac3-98cd-7a4c5be5bb30")
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
