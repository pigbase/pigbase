package main

func main() {
	IsOperatingSystemSupported() // If operating system is not supported, then exit with error
	IsAdmin()                    // If application is not started by admin, then exit with error
	InitDataFile()               // Init data files in Program Files for windows but for linux it's /opt/
	InitWebsocket()			     // Start websocket server
}
