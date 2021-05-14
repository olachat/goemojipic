package goemojipic

import (
	"fmt"
	"testing"
)

func TestApple(t *testing.T) {
	runes := []rune("🌈😃🥰🌍🍞🚗📞🎉🐎🍆🏁🐘🐧🐼")

	for i := 0; i < len(runes); i++ {
		img, err := GetApplePics(runes[i])
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("img loaded %d\n", len(img))
	}
}
