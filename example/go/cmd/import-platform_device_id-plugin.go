package main

// DO NOT EDIT, this file is generated by hover at compile-time for the platform_device_id plugin.

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	platform_device_id "github.com/BestBurning/platform_device_id/go"
)

func init() {
	// Only the init function can be tweaked by plugin maker.
	options = append(options, flutter.AddPlugin(&platform_device_id.PlatformDeviceIdPlugin{}))
}
