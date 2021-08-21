package midiled

import (
	log "github.com/sirupsen/logrus"
)

func BlendColors(input uint32, current uint32) uint32 {
	log.Printf("current led value: %d", current)
	inputRed, inputGreen, inputBlue := IntToRGB(input)
	currentRed, currentGreen, currentBlue := IntToRGB(current)
	log.Printf("current red led value: %d", currentRed)
	log.Printf("current green led value: %d", currentGreen)
	log.Printf("current blue led value: %d", currentBlue)
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
