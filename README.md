# goemojipic

Required `go 1.16` or above; as [embed](https://golang.org/doc/go1.16#library-embed) is used.

`go get https://github.com/olachat/goemojipic` to install

```go
texts := goemojipic.SplitEmojiString("a 🥰 b")
// texts will a slice of EmojiString
// texts[0].Text == "a " texts[0].IsEmoji == false
// texts[1].Text == "🥰" texts[1].IsEmoji == true
// texts[2].Text == " b" texts[2].IsEmoji == false

imgData, err := goemojipic.GetApplePics("🥰")
// imgData: 72x72 png bytes from Apple's font
// err will not be nil if emoji not found

imgData, err := goemojipic.GetGooglePics("🥰")
// imgData: 72x72 png bytes from Google's font
// err will not be nil if emoji not found
```

# memo

* Apple & Google emoji images are extracted by `go run script/main.go`
* emoji parsing using https://github.com/Andrew-M-C/go.emoji

# Reference

* http://www.unicode.org/emoji/charts/full-emoji-list.html
* https://segmentfault.com/a/1190000022100299
* https://github.com/golang/freetype/issues/45
* https://stackoverflow.com/questions/30757193/find-out-if-character-in-string-is-emoji
