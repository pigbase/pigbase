package main

import (
	"log"
	"runtime"
)

func QuitUnsupportedOperatingSystem() {
	log.Fatalf("%s operating system is not supported", runtime.GOOS)
}

func IsOperatingSystemSupported() bool {
	support := false
	switch runtime.GOOS {
	case "windows":
		support = true
		break
	case "linux":
		support = true
		break
	default:
		QuitUnsupportedOperatingSystem()
	}
	return support
}
