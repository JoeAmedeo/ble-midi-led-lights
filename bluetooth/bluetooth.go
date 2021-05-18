package bluetooth

import (
	"fmt"
	"os"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	log "github.com/sirupsen/logrus"
)

func Connect(macAddress string) (chan *bluez.PropertyChanged, error) {
	log.Infoln("Starting connection process")
	log.Infof("MAC Address: %s", macAddress)

	//clean up connection on exit
	defer api.Exit()

	a, err := adapter.GetDefaultAdapter()
	if err != nil {
		return nil, err
	}

	// This command fails if we dont run a scan prior to.
	// have to find a way to do this programatically
	// for now, can use the pi terminal, run `sudo bluetoothclt, agent on, scan on`
	mydevice, err := a.GetDeviceByAddress(macAddress)
	if err != nil {
		return nil, fmt.Errorf("obtaining device failed: %s", err)
	}
	// don't know yet if I'll have to clear the bluez cache, if needed use a.FlushDevices()

	// Have to make sure it's paired first

	log.Infof("device address: %s", mydevice.Properties.Address)
	log.Infof("device name: %s", mydevice.Properties.Name)

	isConnected, err := mydevice.GetConnected()

	if err != nil {
		return nil, fmt.Errorf("getting connection status failed: %s", err)
	}

	if !isConnected {
		log.Infoln("Attempting to connect")
		err = mydevice.Connect()
		if err != nil {
			return nil, fmt.Errorf("connect failed: %s", err)
		}
	}

	log.Infoln("Connection succeeded!")

	changedPropertyChannel, err := mydevice.WatchProperties()

	if err != nil {
		return nil, fmt.Errorf("obtaining variable watch channel failed: %s", err)
	}

	return changedPropertyChannel, nil
}

func KillOnDisconnect(changedPropertyChannel chan *bluez.PropertyChanged, killChannel chan os.Signal, Exit func(int)) {
	for {
		select {
		case property := <-changedPropertyChannel:
			log.Infof("changed property name: %s", property.Name)
			log.Infof("changed property value: %s", property.Value)
			if property.Name == "Connected" && property.Value == false {
				// TODO: attempt to connect to device every few seconds until it succeeds
				log.Infoln("Device was disconnected")
			}
		case <-killChannel:
			Exit(0)
		}
	}
}
