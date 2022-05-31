package main

import (
	"log"
	"os"
)

func IsAdmin() {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		log.Fatal("Application must be started as admin. On Linux just add before command 'sudo' on windows: https://www.windowscentral.com/how-run-app-administrator-windows-10")
	}
	return
}
