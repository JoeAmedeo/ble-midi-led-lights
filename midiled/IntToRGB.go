package midiled

func IntToRGB(rgb uint32) (uint32, uint32, uint32) {
	red := (rgb >> 16) & 0xFF
	green := (rgb >> 8) & 0xFF
	blue := rgb & 0xFF
	return red, green, blue
}
