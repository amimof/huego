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

Include it in your code. You may use `huego.New()` if you've already created a user and know the ip-address/hostname to your bridge.
```Go
package main

import (
  "github.com/amimof/huego"
  "fmt"
)

func main() {
  bridge, err := huego.New("192.168.1.59", "username")
  l, err := bridge.GetLights()
  if err != nil {
    fmt.Fatal(err)
  }
  fmt.Printf("Found %d lights", len(l))
}
```

To discover new bridges and add an user, use `huego.Discover()` and `huego.Login()`
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