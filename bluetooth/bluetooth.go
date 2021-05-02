package bluetooth

import (
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	log "github.com/sirupsen/logrus"
)

func Run(macAddress string) error {

	log.Infoln("Starting connection process")
	log.Infof("MAC Address: %s", macAddress)

	//clean up connection on exit
	defer api.Exit()

	a, err := adapter.GetDefaultAdapter()
	if err != nil {
		return err
	}

	// don't know yet if I'll have to clear the bluez cache, if needed use a.FlushDevices()

	// Have to make sure it's paired first
	drumkit, err := a.GetDeviceByAddress(macAddress)
	if err != nil {
		return err
	}

	log.Infoln(drumkit)

	if drumkit.Properties.Paired == false {
		log.Infoln("Device needs to be paired")
		err := drumkit.Pair()
		if err != nil {
			return err
		}
	}

	if drumkit.Properties.Connected == false {
		log.Infoln("Device needs to be connected")
		err := drumkit.Connect()
		if err != nil {
			return err
		}
	}

	log.Infoln("Connection step complete")

	return nil
}
