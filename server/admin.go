package main

import (
	"log"
	"os"
	"runtime"
)

func IsAdmin() {
	if runtime.GOOS == "windows" {
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		if err != nil {
			// log.Fatal(err)
			log.Fatal("Application must be started as admin: https://www.windowscentral.com/how-run-app-administrator-windows-10")
		}
	} else {
		log.Println("[ WARNING ] Be sure that application is run by administrator, add before command 'sudo'")
	}
	return
}
