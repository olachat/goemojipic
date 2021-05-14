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

// GetApplePic returns 72*72 png bytes from Apple's font if emoji is found
func GetApplePic(emoji string) ([]byte, error) {
	code := getCode(emoji)

	return applePics.ReadFile("apple/" + code + ".png")
}

// GetGooglePic returns 72*72 png bytes from Google's font if emoji is found
func GetGooglePic(emoji string) ([]byte, error) {
	code := getCode(emoji)
	return googlePics.ReadFile("google/" + code + ".png")
}

// EmojiString stores a Text, and indicate if the Text IsEmoji
type EmojiString struct {
	Text    string
	IsEmoji bool
}

// SplitEmojiString split given string emoji & non-emoji string, returning slice of EmojiString
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
