package javbus

import (
	"fmt"
	"testing"
)

func TestGetStarInfo(t *testing.T) {
	starId := "xbm"
	starType := Normal
	info, err := GetStarInfo(starId, starType)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Starinfo: %v", info)
}
