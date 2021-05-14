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
	in, err := midi.Connect()
	if err != nil {
		log.Errorln(err)
	}
	if in != nil {
		defer in.Close()
	}

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go bluetooth.KillOnDisconnect(changedPropertyChannel, sig, os.Exit)

	waitGroup.Wait()
}
