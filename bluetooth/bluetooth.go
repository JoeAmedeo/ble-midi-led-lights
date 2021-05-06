package bluetooth

import (
	"fmt"

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

	mydevice, err := a.GetDeviceByAddress(macAddress)
	if err != nil {
		return fmt.Errorf("obtaining device failed: %s", err)
	}
	// don't know yet if I'll have to clear the bluez cache, if needed use a.FlushDevices()

	// Have to make sure it's paired first

	log.Infof("device address: %s", mydevice.Properties.Address)
	log.Infof("device name: %s", mydevice.Properties.Name)

	isConnected, err := mydevice.GetConnected()

	if err != nil {
		return fmt.Errorf("getting connection status failed: %s", err)
	}

	if !isConnected {
		log.Infoln("Attempting to connect")
		err = mydevice.Connect()
		if err != nil {
			return fmt.Errorf("connect failed: %s", err)
		}
	}

	log.Infoln("Connection succeeded!")

	return nil
}
