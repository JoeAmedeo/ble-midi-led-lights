package midiled

func BlendColors(input uint32, current uint32) uint32 {
	inputRed, inputGreen, inputBlue := IntToRGB(input)
	currentRed, currentGreen, currentBlue := IntToRGB(current)
	red := useInputIfExists(inputRed, currentRed)
	green := useInputIfExists(inputGreen, currentGreen)
	blue := useInputIfExists(inputBlue, currentBlue)
	return RGBToInt(red, green, blue)
}

func useInputIfExists(input uint32, current uint32) uint32 {
	if input > 0 {
		return input
	}
	return current
}
