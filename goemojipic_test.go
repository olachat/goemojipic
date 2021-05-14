package goemojipic

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	data := SplitEmojiString("a 🥰 b🌈😃🥰🌍🍞🚗📞🇨🇮🏴󠁧󠁢󠁷󠁬󠁳󠁿3️⃣🎉🐎🍆🏁🐘🐧🐼")

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

	data = SplitEmojiString("👨‍👩‍👧‍👦")
	if len(data) != 1 && data[0].Text != "👨‍👩‍👧‍👦" && data[0].IsEmoji != true {
		t.Error("Failed to handle: a")
	}

	data = SplitEmojiString("👨‍👩‍👧‍👦1")
	if len(data) != 2 && data[0].Text != "👨‍👩‍👧‍👦" && data[0].IsEmoji != true {
		t.Error("Failed to handle: 👨‍👩‍👧‍👦1")
	}
	if data[1].Text != "1" && data[0].IsEmoji != false {
		t.Error("Failed to handle: 👨‍👩‍👧‍👦1")
	}
}
