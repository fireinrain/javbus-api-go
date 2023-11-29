package javbus

import (
	"javbus-api-go/config"
	"strings"
)

const PAGE_REG = `^\d+$`

var JavbusBaseUrl = config.GlobalConfig.JavbusSite.JavbusUrls[0]

func FormatImageUrl(url string) string {
	if url != "" && !strings.HasPrefix(url, "http") {
		return JavbusBaseUrl + url
	}
	return url
}
