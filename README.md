[![huego](https://godoc.org/github.com/amimof/huego?status.svg)](https://godoc.org/github.com/amimof/huego) [![Go Report Card](https://goreportcard.com/badge/github.com/amimof/huego)](https://goreportcard.com/report/github.com/amimof/huego) [![Coverage](http://gocover.io/_badge/github.com/amimof/huego)](http://gocover.io/github.com/amimof/huego) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go)

# Huego

An extensive Philips Hue client library for [`Go`](https://golang.org/) with an emphasis on simplicity. It is designed to be clean, unbloated and extensible. With `Huego` you can interact with any Philips Hue bridge and its resources including `Lights`, `Groups`, `Scenes`, `Sensors`, `Rules`, `Schedules`, `Resourcelinks`, `Capabilities` and `Configuration` .

![](./logo/logo.png)


## Installation
Get the package and import it in your code.
```
go get github.com/amimof/huego
```
You may use [`New()`](https://godoc.org/github.com/amimof/huego#New) if you have already created an user and know the IP address to your bridge.
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
Or discover a bridge on your network with [`Discover()`](https://godoc.org/github.com/amimof/huego#Discover) and create a new user with [`CreateUser()`](https://godoc.org/github.com/amimof/huego#Bridge.CreateUser).
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

## Contributing

All help in any form is highly appreciated and your are welcome participate in developing `Huego` together. To contribute submit a `Pull Request`. If you want to provide feedback, open up a Github `Issue` or contact me personally. 