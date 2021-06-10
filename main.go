//shows how to watch for new devices and list them
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
	"gitlab.com/gomidi/midi/reader"
	driver "gitlab.com/gomidi/rtmididrv"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

// below two functions are taken from java implementation here: https://stackoverflow.com/questions/4801366/convert-rgb-values-to-integer
func RGBtoInt(red, green, blue uint32) uint32 {
	rgb := red
	rgb = (rgb << 8) + green
	rgb = (rgb << 8) + blue
	return rgb
}

func InttoRGB(rgb uint32) (uint32, uint32, uint32) {
	red := (rgb >> 16) & 0xFF
	green := (rgb >> 8) & 0xFF
	blue := rgb & 0xFF
	return red, green, blue
}

// for now, set all LEDs to a random color
func setAllLeds(device *ws2811.WS2811) error {
	currentColor := RGBtoInt(255, 0, 0)
	for i := 0; i < len(device.Leds(0)); i++ {
		log.Printf("current led: %d", i)
		device.Leds(0)[i] = currentColor
		log.Printf("current led value: %d", device.Leds(0)[i])
	}
	return device.Render()
}

func fade(color uint32, factor float64) uint32 {
	if color == 0 {
		return color
	}
	return uint32(float64(color)*factor) - 1
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

	myReader := reader.New(
		reader.NoteOn(func(p *reader.Position, channel, key, velocity uint8) {
			err := setAllLeds(device)
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

	fadeConstant := 0.90

	go func() {
		for {
			for i := 0; i < len(device.Leds(0)); i++ {
				red, green, blue := InttoRGB(device.Leds(0)[i])
				log.Infof("current RGB values: r -> %s, g -> %s, b -> %s", red, green, blue)
				device.Leds(0)[i] = RGBtoInt(fade(red, fadeConstant), fade(green, fadeConstant), fade(blue, fadeConstant))
			}
			err := device.Render()
			if err != nil {
				log.Errorf(`failed to dim lights: %s`, err)
			}
			err = device.Wait()
			if err != nil {
				log.Errorf(`failed to wait for render lights: %s`, err)
			}
		}
	}()

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
