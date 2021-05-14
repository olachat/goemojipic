package goemojipic

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/Andrew-M-C/go.emoji/official"
)

//go:embed apple/*
var applePics embed.FS

//go:embed google/*
var googlePics embed.FS

func getCode(text string) string {
	var code string
	for _, c := range text {
		code += fmt.Sprintf("%U", c)
	}
	return code
}

func GetApplePics(emoji string) ([]byte, error) {
	code := getCode(emoji)

	return applePics.ReadFile("apple/" + code + ".png")
}

func GetGooglePics(emoji string) ([]byte, error) {
	code := getCode(emoji)
	return googlePics.ReadFile("google/" + code + ".png")
}

type EmojiString struct {
	Text    string
	IsEmoji bool
}

func SplitEmojiString(s string) (result []EmojiString) {
	buff := bytes.Buffer{}
	nextIndex := 0

	for i, r := range s {
		if i < nextIndex {
			continue
		}

		match, length := official.AllSequences.HasEmojiPrefix(s[i:])
		if false == match {
			buff.WriteRune(r)
			continue
		}

		nextIndex = i + length
		text := buff.String()
		buff.Reset()
		if text != "" {
			result = append(result, EmojiString{text, false})
		}

		result = append(result, EmojiString{s[i : i+length], true})
	}

	text := buff.String()
	if text != "" {
		result = append(result, EmojiString{text, false})
	}

	return
}
