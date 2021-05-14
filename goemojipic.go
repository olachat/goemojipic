package goemojipic

import "embed"

//go:embed apple/*
var applePics embed.FS

//go:embed google/*
var googlePics embed.FS

func GetApplePics(emoji string) ([]byte, error) {
	return applePics.ReadFile("apple/" + emoji + ".png")
}

func GetGooglePics(emoji string) ([]byte, error) {
	return googlePics.ReadFile("google/" + emoji + ".png")
}
