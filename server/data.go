package main

import (
	"errors"
	"log"
	"os"
	"runtime"
)

func PathOfFile() string {
	os := runtime.GOOS
	outputPath := ""
	switch os {
	case "windows":
		outputPath = "C:\\Program Files\\Pigbase\\data"
		outputPath = "operatingSystemIsNotSupported"
		break
	case "linux":
		outputPath = "/opt/pigbase/data"
		break
	case "darwin":
		outputPath = "operatingSystemIsNotSupported"
		break
	default:
		outputPath = "operatingSystemIsNotSupported"
		break
	}
	if outputPath == "operatingSystemIsNotSupported" {
		log.Fatalf("You're actuall operating system is not supported. Actual OS: %s\n", os)
	}
	return outputPath
}

func InitData() {
	if _, err := os.Stat("./data.encoded"); err == nil {

	} else if errors.Is(err, os.ErrNotExist) {

	} else {

	}
}
