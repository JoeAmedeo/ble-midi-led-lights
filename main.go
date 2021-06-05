//shows how to watch for new devices and list them
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
	"gitlab.com/gomidi/midi/reader"
	driver "gitlab.com/gomidi/rtmididrv"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func randomUInt32(min, max uint32) uint32 {
	var a = rand.Uint32()
	a %= (max - min)
	a += min
	return a
}

// for now, set all LEDs to a random color
func setAllLeds(device *ws2811.WS2811, on bool) error {
	onColor := uint32(0xffffff)
	offColor := uint32(0x000000)
	log.Printf("color values: %d, %d", onColor, offColor)
	for i := 0; i < len(device.Leds(0)); i++ {
		log.Printf("current led: %d", i)
		var currentColor uint32
		if on {
			currentColor = onColor
		} else {
			currentColor = offColor
		}
		log.Printf("current on status: %t", on)
		device.Leds(0)[i] = currentColor
		log.Printf("current led value: %d", device.Leds(0)[i])
	}
	return device.Render()
}

func main() {

	myDriver, err := driver.New()

	if err != nil {
		panic(fmt.Errorf("creating driver failed: %s", err))
	}

	defer myDriver.Close()

	ins, err := myDriver.Ins()

	if err != nil {
		panic(fmt.Errorf("getting input channels failed: %s", err))
	}

	in := ins[1]

	defer in.Close()

	err = in.Open()

	if err != nil {
		panic(fmt.Errorf("openning input failed: %s", err))
	}

	ledOptions := ws2811.DefaultOptions
	ledOptions.Channels[0].LedCount = 4

	device, err := ws2811.MakeWS2811(&ledOptions)

	if err != nil {
		panic(fmt.Errorf("failed to create LED device: %s", err))
	}

	err = device.Init()

	if err != nil {
		panic(fmt.Errorf("failed to initialize LED device: %s", err))
	}

	defer device.Fini()

	on := false

	myReader := reader.New(
		reader.NoteOn(func(p *reader.Position, channel, key, velocity uint8) {
			err := setAllLeds(device, on)
			on = !on
			if err != nil {
				log.Printf("error rendering lights: %s", err)
			}
		}),
	)

	err = myReader.ListenTo(in)

	if err != nil {
		panic(fmt.Errorf("reading from input failed: %s", err))
	}

	log.Println("Midi listener added without errors!")

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func(killChannel chan os.Signal, Exit func(int)) {
		for {
			select {
			case <-killChannel:
				Exit(0)
			}
		}
	}(sig, os.Exit)

	waitGroup.Wait()
}
