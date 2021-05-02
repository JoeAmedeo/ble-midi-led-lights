package bluetooth

import (
	"fmt"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/agent"
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
	devices, err := a.GetDevices()
	if err != nil {
		return err
	}

	found := false

	for _, device := range devices {

		if device.Properties.Address != macAddress {
			continue
		}

		if device.Properties.Paired {
			continue
		}

		found = true
		// log.Info(i, v.Path)
		log.Infof("Pairing with %s", device.Properties.Address)

		err := device.Pair()
		if err != nil {
			return fmt.Errorf("pair failed: %s", err)
		}

		log.Info("Pair succeed, connecting...")

		agent.SetTrusted(adapterId, device.Path())

		err = device.Connect()
		if err != nil {
			return fmt.Errorf("connect failed: %s", err)
		}
	}

	if !found {
		return fmt.Errorf("device not found")
	}

	return nil
}
