package midiled

type LedRange struct {
	Start int
	End   int
}

type KeyColor struct {
	Red   uint32
	Green uint32
	Blue  uint32
	Range LedRange
}

// constants for default midi values from the TD-17 module: https://rolandus.zendesk.com/hc/en-us/articles/360005173411-TD-17-Default-Factory-MIDI-Note-Map
const (
	KICK                uint8 = 36
	SNARE_HEAD          uint8 = 38
	SNARE_RIM           uint8 = 40
	SNARE_X_STICK       uint8 = 37
	TOM_1_HEAD          uint8 = 48
	TOM_1_RIM           uint8 = 50
	TOM_2_HEAD          uint8 = 45
	TOM_2_RIM           uint8 = 47
	TOM_3_HEAD          uint8 = 43
	TOM_3_RIM           uint8 = 58
	HIGHHAT_OPEN_BOW    uint8 = 46
	HIGHHAT_OPEN_EDGE   uint8 = 26
	HIGHHAT_CLOSED_BOW  uint8 = 42
	HIGHHAT_CLOSED_EDGE uint8 = 22
	HIGHHAT_PEDAL       uint8 = 44
	CRASH_1_BOW         uint8 = 49
	CRASH_1_EDGE        uint8 = 55
	CRASH_2_BOW         uint8 = 57
	CRASH_2_EDGE        uint8 = 52
	RIDE_BOW            uint8 = 51
	RIDE_EDGE           uint8 = 59
	RIDE_BELL           uint8 = 53
	AUX_HEAD            uint8 = 27
	AUX_RIM             uint8 = 28
)

const TOTAL_LEDS = 155

var SNARE_LEDS = LedRange{0, 52}
var TOM_1_LEDS = LedRange{53, 86}
var TOM_2_LEDS = LedRange{87, 120}
var TOM_3_LEDS = LedRange{121, 154}
var ALL_LEDS = LedRange{0, 154}
var NO_LEDS = LedRange{155, 155}
