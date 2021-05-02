//shows how to watch for new devices and list them
package main

import (
	"os"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	log "github.com/sirupsen/logrus"
)

func Run(macAddress string) error {

	log.Infoln("Starting connection process")

	//clean up connection on exit
	defer api.Exit()

	a, err := adapter.GetDefaultAdapter()
	if err != nil {
		return err
	}

	// don't know yet if I'll have to clear the bluez cache, if needed use a.FlushDevices()

	// Have to make sure it's paired first
	drumkit, err := a.GetDeviceByAddress("TODO: parameterize the mac address")
	if err != nil {
		return err
	}

	if !drumkit.Properties.Paired {
		log.Infoln("Device needs to be paired")
		err := drumkit.Pair()
		if err != nil {
			return err
		}
	}

	if !drumkit.Properties.Connected {
		log.Infoln("Device needs to be connected")
		err := drumkit.Connect()
		if err != nil {
			return err
		}
	}

	log.Infoln("Connection step complete")

	return nil
}

func main() {
	macAddress := os.Args[0]
	err := Run(macAddress)
	if err != nil {
		log.Errorln(err)
	}
}
