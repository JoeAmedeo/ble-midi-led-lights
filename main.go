//shows how to watch for new devices and list them
package main

import (
	"ble-midi-drums/bluetooth"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	macAddress := os.Args[0]
	err := bluetooth.Run(macAddress)
	if err != nil {
		log.Errorln(err)
	}
}
