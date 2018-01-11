# Huego

An extensive Philips Hue client library for [`Go`](https://golang.org/) with an emphasis on simplicity.

This project is currently in **ALPHA** and not recommended for production use. All help in any form is highly appreciated. You are more than welcome to contact me if you have feedback, feature requests, report bugs etc.

## Installation
Get the package
```
go get github.com/amimof/huego
```

Include it in your code
```Go
package main

import (
  "github.com/amimof/huego"
  "fmt"
)

func main() {
  bridge, err := huego.New("username", "password")
  l, err := hue.GetLights()
  if err != nil {
    fmt.Fatal(err)
  }
  fmt.Printf("Found %d lights", len(l))
}
```

## Project Status

This project is currently in **ALPHA** and still under heavy development. Current iteration is subject to big changes until the initial release. Below is the current status of *modules* that are expected to be implemented.

| Module | Functions | Tests |
| ------ | ------ | ------ |
| Lights | `Complete` | `Complete` |
| Groups | `Complete` | `Complete` |
| Sensors | `Complete` | `Complete` |
| Schedules | `Complete` | `Complete` |
| Scenes | `Complete` | `Complete` |
| Rules | `Complete` | `Complete` |
| Resourcelinks | `Complete` | `Complete` |
| Configuration | `Complete`  | `Complete` 
| Capabilities | `Not Started` | `Not Started` 

Other than above core modules, each module needs additional *helper* methods for conveniance and flavour. The goal is to keep it simple, and not to bloat the library with functionality that developers might want to write on their own. 

## Goal

The goal of this project is to provide an easy to use, stable and extensive library that is up to spec with the official [Philips Hue API](https://www.developers.meethue.com/philips-hue-api). It should be possible to interact with *all* API endpoints that is available on a Philips Hue bridge through the package(s) in this repository.

## To-Do

* Add helper methods on each module
* Add `SetSceneLightState`
* Add `RecallScene`
* Finish `Capabilities`
* More tests