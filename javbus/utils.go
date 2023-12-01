package javbus

import (
	"fmt"
	"image"
	"math/rand"
	"time"

	// Required to decode GIF images
	_ "image/gif"
	// Required to decode JPEG images
	_ "image/jpeg"
	// Required to decode PNG images
	_ "image/png"
	"javbus-api-go/config"
	"net/http"
	"strings"
)

const PAGE_REG = `^\d+$`

var JavbusBaseUrl = config.GlobalConfig.JavbusSite.JavbusUrls[0]
var DefaultRandomSource = rand.NewSource(time.Now().UnixNano())
var RandomGen = rand.New(DefaultRandomSource)

// FormatImageUrl
//
//	@Description: 格式化图片url
//	@param url
//	@return string
func FormatImageUrl(url string) string {
	if url != "" && !strings.HasPrefix(url, "http") {
		return JavbusBaseUrl + url
	}
	return url
}

// GetImageSize
//
//	@Description: 获取给定图片链接的图片宽高信息
//	@param imageUrl
//	@return ImageSize
//	@return error
func GetImageSize(imageUrl string) (ImageSize, error) {
	imageSize := ImageSize{}
	// Fetch the image
	resp, err := http.Get(imageUrl)
	if err != nil {
		//fmt.Println("Error fetching image:", err)
		return imageSize, err
	}
	defer resp.Body.Close()

	// Decode the image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		//fmt.Println("Error decoding image:", err)
		return imageSize, fmt.Errorf("error decoding image: %v", err)
	}
	// Get the width and height of the image
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	//fmt.Printf("Original Width: %d, Height: %d\n", width, height)

	imageSize.Height = height
	imageSize.Width = width

	return imageSize, nil
}

func RandomCookieHeader() map[string]string {
	m := make(map[string]string, 2)
	cookies := config.GlobalConfig.JavbusSite.UserCookies
	userAgents := config.GlobalConfig.JavbusSite.UserAgents
	// Randomly choose an element from the list
	randomIndex := RandomGen.Intn(len(cookies))
	randomCookie := cookies[randomIndex]
	m["Cookie"] = randomCookie.Cookies

	randomAgent := RandomGen.Intn(len(userAgents))
	randomAgentStr := userAgents[randomAgent]
	m["User-Agent"] = randomAgentStr
	return m
}
