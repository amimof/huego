package main

import (
	"context"
	"fmt"
	"net/http"
	//"crypto/tls"
	"github.com/amimof/huego/v2"
	"encoding/json"
	"io/ioutil"
)

func errorHandler(res *http.Response) error {
	var a *huego.APIResponse
	if res.StatusCode >= 400 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, a)
		if err != nil {
			return err
		}
		for _, e := range a.Errors {
			return e
		}
	}
	return nil
}

func main() {

	// Create the client
	c, err := huego.NewInsecureClientV2("192.168.12.11", "GAEe2hlIhYnAsFdZIVD5OsGKXkzG7cz3hmIEqbhi")
	if err != nil {
		panic(err)
	}

	// // Get a light with custom client
	// res, err := 
	// 	huego.NewRequest(c.Clip).
	// 		Verb(http.MethodGet).
	// 		Resource(huego.TypeLight).
	// 		//ID("b1312d43-7aca-447b-92eb-95402aace153").
	// 		OnError(errorHandler).
	// 		Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// //fmt.Printf("%s\n", string(res.BodyRaw))
	// var re huego.APIResponse
	// err = res.Into(&re)
	// if err != nil {
	// 	panic(err)
	// }
	// var lights []huego.Light
	// err = re.Into(&lights) 
	// if err != nil {
	// 	panic(err)
	// }
	// for _, light := range lights {
	// 	fmt.Printf("Light: %s\n", *light.ID)
	// }


	// ures, err := 
	// 	huego.NewRequest(c.Clip).
	// 		Verb(http.MethodPut).
	// 		Resource(huego.TypeLight).
	// 		ID("b1312d43-7aca-447b-92eb-95402aace153").
	// 		Body(res.BodyRaw).
	// 		OnError(errorHandler).
	// 		Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Res: %s\n", string(ures.BodyRaw))


	//Get one light
	// light, err := c.GetLight("b1312d43-7aca-447b-92eb-95402aace153")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("ID: %s\nIDv1: %s\n", *light.ID, *light.IDv1)
	// fmt.Printf("On: %t\n", *light.On.On)

	// // Update the light
	// ison := false
	// light.On.On = &ison
	// res, err := c.SetLightContext(context.Background(), *light.ID, light)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Response: %s\n", string(res.BodyRaw))

	//Get many lights
	lights, err := c.GetLightsContext(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found %d lights\n", len(lights))
	for i, light := range lights {
		fmt.Printf("%d: %s\n", i, *light.ID)
	}
}