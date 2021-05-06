package bluetooth

import (
	"fmt"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/agent"
	"github.com/muka/go-bluetooth/bluez/profile/device"
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

	adapterId, err := a.GetAdapterID()
	if err != nil {
		return err
	}

	log.Infof("Adapter created: %s", a.Properties.Name)
	// don't know yet if I'll have to clear the bluez cache, if needed use a.FlushDevices()

	// Have to make sure it's paired first

	discovery, cancel, err := api.Discover(a, nil)

	if err != nil {
		return fmt.Errorf("discovery failed: %s", err)
	}

	defer cancel()

	found := false

	for discoveredDevice := range discovery {

		mydevice, err := device.NewDevice1(discoveredDevice.Path)

		log.Infof("device path: %s", discoveredDevice.Path)
		log.Infof("device address: %s", mydevice.Properties.Address)
		log.Infof("device name: %s", mydevice.Properties.Name)

		if err != nil {
			return fmt.Errorf("creating device failed: %s", err)
		}

		if mydevice.Properties.Address != macAddress {
			continue
		}

		if mydevice.Properties.Paired {
			continue
		}

		found = true
		// log.Info(i, v.Path)
		log.Infof("Pairing with %s", mydevice.Properties.Address)

		err = mydevice.Pair()

		if err != nil {
			return fmt.Errorf("pair failed: %s", err)
		}

		log.Info("Pair succeed, connecting...")

		agent.SetTrusted(adapterId, mydevice.Path())

		err = mydevice.Connect()
		if err != nil {
			return fmt.Errorf("connect failed: %s", err)
		}
	}

	if !found {
		return fmt.Errorf("device not found")
	}

	log.Infoln("Connection succeeded!")

	return nil
}
