package bluetooth

import (
	"fmt"

	"github.com/godbus/dbus/v5"
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

	//Connect DBus System bus
	conn, err := dbus.SystemBus()
	if err != nil {
		return err
	}

	ag := agent.NewSimpleAgent()
	err = agent.ExposeAgent(conn, ag, agent.CapKeyboardDisplay, true)
	if err != nil {
		return fmt.Errorf("SimpleAgent: %s", err)
	}

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

		log.Infof("%s --- %s", device.Properties.Name, device.Properties.Address)

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

	log.Infoln("Connection succeeded!")

	return nil
}
