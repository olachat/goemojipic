package goemojipic

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	data := SplitEmojiString("a ๐ฅฐ b๐๐๐ฅฐ๐๐๐๐๐จ๐ฎ๐ด๓ ง๓ ข๓ ท๓ ฌ๓ ณ๓ ฟ3๏ธโฃ๐๐๐๐๐๐ง๐ผ")

	for _, str := range data {
		if str.IsEmoji {
			img, err := GetApplePic(str.Text)
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("img loaded %d\n", len(img))
		}
	}

	data = SplitEmojiString("")
	if len(data) != 0 {
		t.Error("Failed to handle empty string")
	}

	data = SplitEmojiString("a")
	if len(data) != 1 && data[0].Text != "a" && data[0].IsEmoji != false {
		t.Error("Failed to handle: a")
	}

	data = SplitEmojiString("๐จโ๐ฉโ๐งโ๐ฆ")
	if len(data) != 1 && data[0].Text != "๐จโ๐ฉโ๐งโ๐ฆ" && data[0].IsEmoji != true {
		t.Error("Failed to handle: a")
	}

	data = SplitEmojiString("๐จโ๐ฉโ๐งโ๐ฆ1")
	if len(data) != 2 && data[0].Text != "๐จโ๐ฉโ๐งโ๐ฆ" && data[0].IsEmoji != true {
		t.Error("Failed to handle: ๐จโ๐ฉโ๐งโ๐ฆ1")
	}
	if data[1].Text != "1" && data[0].IsEmoji != false {
		t.Error("Failed to handle: ๐จโ๐ฉโ๐งโ๐ฆ1")
	}
}
