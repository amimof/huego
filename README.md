[![huego](https://godoc.org/github.com/amimof/huego?status.svg)](https://godoc.org/github.com/amimof/huego) 
[![Go Report Card](https://goreportcard.com/badge/github.com/amimof/huego)](https://goreportcard.com/report/github.com/amimof/huego)
[![Coverage](http://gocover.io/_badge/github.com/amimof/huego)](http://gocover.io/github.com/amimof/huego)

# Huego

An extensive Philips Hue client library for [`Go`](https://golang.org/) with an emphasis on simplicity. It is designed to be clean, unbloated and extensible. With `Huego` you can interact with any Philips Hue bridge and its resources including `Lights`, `Groups`, `Scenes`, `Sensors`, `Rules`, `Schedules`, `Resourcelinks`, `Capabilities` and `Configuration` .

![](./logo/logo.png)


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

See [godoc.org/github.com/amimof/huego](https://godoc.org/github.com/amimof/huego) for the full package documentation.

## Testing

The tests requires an accessible Philips Hue Bridge IP address and a pre-configured username for authenticating. Before running the tests, make sure to set the environment variables `HUE_HOSTNAME` and `HUE_USERNAME`. If you don't have an username, you may create one using [`CreateUser()`](https://godoc.org/github.com/amimof/huego#Bridge.CreateUser) or refer to the official [Getting Started Guide](https://www.developers.meethue.com/documentation/getting-started).

## Contributing

All help in any form is highly appreciated and your are welcome participate in developing `Huego` together. To contribute, clone the `master` branch and create a `Pull Request`. If you want provide feedback, open up a `New Issue` or contact me personally. 