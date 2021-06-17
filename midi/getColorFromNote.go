package midi

func getColorFromNote(key uint8, velocity uint8) (uint32, uint32, uint32, LedRange) {
	switch key {
	case KICK:
		return 0, 0, uint32(velocity), ALL_LEDS

	case SNARE_HEAD, SNARE_RIM, SNARE_X_STICK:
		return uint32(velocity), 0, 0, SNARE_LEDS

	case TOM_1_HEAD, TOM_1_RIM:
		return uint32(velocity), 0, 0, TOM_1_LEDS

	case TOM_2_HEAD, TOM_2_RIM:
		return uint32(velocity), 0, 0, TOM_2_LEDS

	case TOM_3_HEAD, TOM_3_RIM:
		return uint32(velocity), 0, 0, TOM_3_LEDS

	case HIGHHAT_CLOSED_BOW, HIGHHAT_CLOSED_EDGE, HIGHHAT_OPEN_BOW, HIGHHAT_OPEN_EDGE, HIGHHAT_PEDAL:
		return 0, uint32(velocity), 0, SNARE_LEDS

	case CRASH_1_BOW, CRASH_1_EDGE:
		return 0, uint32(velocity), 0, TOM_1_LEDS

	case RIDE_BELL, RIDE_BOW, RIDE_EDGE:
		return 0, uint32(velocity), 0, TOM_2_LEDS

	case CRASH_2_BOW, CRASH_2_EDGE:
		return 0, uint32(velocity), 0, TOM_3_LEDS

	default:
		return 0, 0, 0, ALL_LEDS
	}
}
