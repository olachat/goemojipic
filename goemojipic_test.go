package goemojipic

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	data := SplitEmojiString("a 🥰 b🌈😃🥰🌍🍞🚗📞🇨🇮🏴󠁧󠁢󠁷󠁬󠁳󠁿3️⃣🎉🐎🍆🏁🐘🐧🐼")

	for _, str := range data {
		if str.IsEmoji {
			img, err := GetApplePics(str.Text)
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("img loaded %d\n", len(img))
		}
	}
}
