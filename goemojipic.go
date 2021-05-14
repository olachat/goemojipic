package goemojipic

import (
	"embed"
	"fmt"
)

//go:embed apple/*
var applePics embed.FS

//go:embed google/*
var googlePics embed.FS

func GetApplePics(emoji rune) ([]byte, error) {
	code := fmt.Sprintf("%U", emoji)
	return applePics.ReadFile("apple/" + code + ".png")
}

func GetGooglePics(emoji rune) ([]byte, error) {
	code := fmt.Sprintf("%U", emoji)
	return googlePics.ReadFile("google/" + code + ".png")
}
