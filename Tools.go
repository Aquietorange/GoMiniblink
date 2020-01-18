package miniblink

import (
	uuid "github.com/satori/go.uuid"
	"image/color"
	"strings"
)

func NewUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func IntToColor(rgba int) color.Color {
	return color.RGBA{
		R: uint8(rgba),
		G: uint8(rgba >> 8),
		B: uint8(rgba >> 16),
		A: uint8(rgba >> 24),
	}
}
