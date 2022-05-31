package main

import (
	"log"
	"os"
	"runtime"
)

func InitDataFile() {
	log.Println("Initializing data file")
	switch runtime.GOOS {
	case "windows":
		if _, err := os.Stat("C:\\Program Files\\Pigbase\\data"); err == nil {
			// File and directory exists so we don't need to create it again.
			log.Println("Data file is already created")
			break
		} else {
			// File and directory doesn't exist so we need to create it.
			err := os.Mkdir("C:\\Program Files\\Pigbase\\", 0755)
			if err != nil {
				log.Fatal(err)
			} else {
				file, err := os.Create("C:\\Program Files\\Pigbase\\data")
				if err != nil {
					log.Fatal(err)
				}
				log.Println("Data file and directory has been created")
				file.Close()
			}
		}
		break
	case "linux":
		// "/opt/pigbase/data"
		if _, err := os.Stat("/opt/pigbase/data"); err == nil {
			// File and directory exists so we don't need to create it again.
			log.Println("Data file is already created")
			break
		} else {
			// File and directory doesn't exist so we need to create it.
			err := os.Mkdir("/opt/pigbase/", 0755)
			if err != nil {
				log.Fatal(err)
			} else {
				file, err := os.Create("/opt/pigbase/data")
				if err != nil {
					log.Fatal(err)
				}
				log.Println("Data file and directory has been created")
				file.Close()
			}
		}
		break
	}
}
