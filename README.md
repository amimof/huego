[![huego](https://godoc.org/github.com/amimof/huego?status.svg)](https://godoc.org/github.com/amimof/huego) 
[![Go Report Card](https://goreportcard.com/badge/github.com/amimof/huego)](https://goreportcard.com/report/github.com/amimof/huego)

# Huego

An extensive Philips Hue client library for [`Go`](https://golang.org/) with an emphasis on simplicity.

![](./logo.png)

_This project is currently in **BETA** and not recommended for production use. All help in any form is highly appreciated. You are more than welcome to contact me if you have feedback, feature requests, report bugs etc._

Ses [godoc.org/github.com/amimof/huego](https://godoc.org/github.com/amimof/huego) for the full package documentation.

## Installation
Get the package
```
go get github.com/amimof/huego
```

Include it in your code. You may use [`New()`](https://godoc.org/github.com/amimof/huego#New) if you have already created an user and know the IP address to your bridge.
```Go
package main

import (
  "github.com/amimof/huego"
  "fmt"
)

func main() {
  bridge := huego.New("192.168.1.59", "username")
  l, err := bridge.GetLights()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Found %d lights", len(l))
}
```

Discover a bridge on your network with [`Discover()`](https://godoc.org/github.com/amimof/huego#Discover) and create a new user with [`CreateUser()`](https://godoc.org/github.com/amimof/huego#Bridge.CreateUser).
```Go
func main() {
  bridge, _ := huego.Discover()
  user, _ := bridge.CreateUser("my awesome hue app") // Link button needs to be pressed
  bridge = bridge.Login(user)
  light, _ := bridge.GetLight(3)
  light.Off()
}
``` 

## Documentation

See [godoc.org/github.com/amimof/huego](https://godoc.org/github.com/amimof/huego)

## Testing
The tests requires an accessible Philips Hue Bridge IP address and a pre-configured username for authenticating. Before running the tests, make sure to set the environment variables `HUE_HOSTNAME` and `HUE_USERNAME`. If you don't have an username, you may create one using [`CreateUser()`](https://godoc.org/github.com/amimof/huego#Bridge.CreateUser) or refer to the official [Getting Started Guide](https://www.developers.meethue.com/documentation/getting-started).