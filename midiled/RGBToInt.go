package midiled

// below two functions are taken from java implementation here: https://stackoverflow.com/questions/4801366/convert-rgb-values-to-integer
func RGBToInt(red, green, blue uint32) uint32 {
	rgb := red
	rgb = (rgb << 8) + green
	rgb = (rgb << 8) + blue
	return rgb
}
