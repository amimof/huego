# Changelog

## 1.0.2
* Improved test coverage
* Use of `httpmock` in tests

## 1.0.1
* Added `go.mod` for Go 1.11 module compatibility

## 1,0.0
Exiting beta and entering stable release

## 1.0.0-beta.2
Much better error handling. Whenever the bridge API returns an error, huego will return that to the user accordingly through an error struct rather than just throwing a json.UnmarshalTypeError.

## 1.0.0-beta.1
* `Sensor.State` and `Sensor.Config` is now `interface{}` since these varry depending on the type
* As of this release `Gitflow` is obsolete.

## 1.0.0-alpha.5
* Added `SetState` receivers to `Group` and `Light`.
* Renamed `SetLight` to `SetLightState` for a more consistent naming convention.
* Implemented `Capabilities`

## 1.0.0-alpha.4
* Huego now has a logo
* Changes to fulfill `golint`, `govet` and `gofmt`

## 1.0.0-alpha.3
* Added `Group` receivers: `Alert()`, `Bri()`, `Ct()`, `Effect()`, `Hue()`, `IsOn()`, `Off()`, `On()`, `Rename()`, `Sat()`, `Scene()`, `TransitionTime()` and `Xy()`
* Added `Light` receivers: `Alert()`, `Effect()` and `TransitionTime()`
* Renamed type `Action` to `State` that belongs to `Group`

## 1.0.0
Initial release
