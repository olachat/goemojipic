package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	f, _ := os.Open("script/full-emoji-list.html")
	defer f.Close()

	doc, _ := goquery.NewDocumentFromReader(f)

	doc.Find("td.code").Each(func(i int, s *goquery.Selection) {
		code := strings.ReplaceAll(s.Find("a").Text(), " ", "")

		img, _ := s.Next().Next().Find("img").Attr("src")
		if len(img) < 100 {
			fmt.Printf("Apple emoji not found: %s %s\n", code, s.Next().Text())
		} else {
			data, _ := base64.StdEncoding.DecodeString(img[22:])
			ioutil.WriteFile("apple/"+code+".png", data, 0644)
		}

		img, _ = s.Next().Next().Next().Find("img").Attr("src")
		if len(img) < 100 {
			fmt.Printf("Google emoji not found: %s %s\n", code, s.Next().Text())
		} else {
			data, _ := base64.StdEncoding.DecodeString(img[22:])
			ioutil.WriteFile("google/"+code+".png", data, 0644)
		}
	})
}
