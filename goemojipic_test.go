package goemojipic

import (
	"fmt"
	"testing"
)

func TestApple(t *testing.T) {
	img, err := GetApplePics("U+00A9")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("img loaded %d", len(img))
}
