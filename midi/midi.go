package midi

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	driver "gitlab.com/gomidi/rtmididrv"
)

// constants for default midi values from the TD-17 module: https://rolandus.zendesk.com/hc/en-us/articles/360005173411-TD-17-Default-Factory-MIDI-Note-Map
const (
	KICK                int = 36
	SNARE_HEAD          int = 38
	SNARE_RIM           int = 40
	SNARE_X_STICK       int = 37
	TOM_1_HEAD          int = 48
	TOM_1_RIM           int = 50
	TOM_2_HEAD          int = 45
	TOM_2_RIM           int = 47
	TOM_3_HEAD          int = 43
	TOM_3_RIM           int = 58
	HIGHHAT_OPEN_BOW    int = 46
	HIGHHAT_OPEN_EDGE   int = 43
	HIGHHAT_CLOSED_BOW  int = 42
	HIGHHAT_CLOSED_EDGE int = 22
	HIGHHAT_PEDAL       int = 44
	CRASH_1_BOW         int = 49
	CRASH_1_EDGE        int = 55
	CRASH_2_BOW         int = 57
	CRASH_2_EDGE        int = 52
	RIDE_BOW            int = 51
	RIDE_EDGE           int = 59
	RIDE_BELL           int = 53
	AUX_HEAD            int = 27
	AUX_RIM             int = 28
)

func Connect() (midi.In, error) {
	myDriver, err := driver.New()

	if err != nil {
		return nil, fmt.Errorf("creating driver failed: %s", err)
	}

	ins, err := myDriver.Ins()
	if err != nil {
		return nil, fmt.Errorf("get input streams failed: %s", err)
	}

	for _, input := range ins {
		log.Printf("input device info: %s", input.String())
	}

	in := ins[0]

	err = in.Open()

	if err != nil {
		return nil, fmt.Errorf("opening input failed: %s", err)
	}

	myReader := reader.New(reader.NoLogger(), reader.Each(func(pos *reader.Position, msg midi.Message) {
		// TODO, This function will trigger
		log.Printf("got message %s\n", msg)
	}),
	)

	err = myReader.ListenTo(in)

	if err != nil {
		return in, fmt.Errorf("reading from input failed: %s", err)
	}

	log.Println("Midi listener added without errors!")

	return in, nil
}
