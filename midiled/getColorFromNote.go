package midiled

func GetColorFromNote(key uint8, velocity uint8) KeyColor {

	keyMap := GetMidiMap()

	keyColor, ok := keyMap[key]

	if ok {
		keyColor.Red *= uint32(velocity)
		keyColor.Green *= uint32(velocity)
		keyColor.Blue *= uint32(velocity)
		return keyColor
	}

	return KeyColor{0, 0, 0, NO_LEDS}
}
