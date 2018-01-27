# Changelog

## 1.0.0-alpha.5
* Added `SetState` receivers to `Group` and `Light`.
* Renamed `SetLight` to `SetLightState` for a more consistent naming convention.

## 1.0.0-alpha.4
* Huego now has a logo
* Changes to fulfill `golint`, `govet` and `gofmt`

## 1.0.0-alpha.3
* Added `Group` receivers: `Alert()`, `Bri()`, `Ct()`, `Effect()`, `Hue()`, `IsOn()`, `Off()`, `On()`, `Rename()`, `Sat()`, `Scene()`, `TransitionTime()` and `Xy()`
* Added `Light` receivers: `Alert()`, `Effect()` and `TransitionTime()`
* Renamed type `Action` to `State` that belongs to `Group`

## 1.0.0
Initial release
