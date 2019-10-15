package main

import "image/color"

const (
	tileSize   = 20
	tileMargin = 2
)

var (
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

func tileBackgroundColor(isOpen bool) color.Color {
	if isOpen {
		return color.RGBA{0xee, 0xe4, 0xda, 0xff}
	}
	return color.NRGBA{0xee, 0xe4, 0xda, 0x59}
}

func tileColor() color.Color {
	return color.RGBA{0x77, 0x6e, 0x65, 0xff}
}
