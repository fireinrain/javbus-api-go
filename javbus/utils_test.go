package javbus

import (
	"fmt"
	"testing"
)

func TestGetImageSize(t *testing.T) {
	var imageUrl string = "https://www.javbus.com/pics/cover/8hqb_b.jpg"
	size, err := GetImageSize(imageUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", size)
}

func TestRandomCookieHeader(t *testing.T) {
	header := RandomCookieHeader()
	fmt.Printf("%v\n", header)
}
