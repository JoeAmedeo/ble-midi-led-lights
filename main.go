//shows how to watch for new devices and list them
package main

import (
	"ble-midi-drums/bluetooth"
	"ble-midi-drums/midi"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	macAddress := os.Args[1]
	changedPropertyChannel, err := bluetooth.Connect(macAddress)
	if err != nil {
		log.Errorln(err)
	}
	err = midi.Connect()
	if err != nil {
		log.Errorln(err)
	}

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		for {
			select {
			case property := <-changedPropertyChannel:
				log.Infof("changed property name: %s", property.Name)
				log.Infof("changed property value: %s", property.Value)
				if property.Name == "Connected" && property.Value == false {
					// TODO: attempt to connect to device every few seconds until it succeeds
					log.Infoln("Device was disconnected")
				}
			case <-sig:
				os.Exit(0)
			}
		}
	}()

	waitGroup.Wait()
}
