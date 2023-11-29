package javbus

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/carlmjohnson/requests"
	"golang.org/x/net/proxy"
	"javbus-api-go/config"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var HttpClient *requests.Builder

var DefaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"

func NewHttpClient() (*requests.Builder, error) {
	requestTimeout := config.GlobalConfig.JavbusSite.RequestTimeout
	requestTimeoutSecond := time.Duration(requestTimeout) * time.Second
	//检查是否开启proxy
	var client *http.Client
	if config.GlobalConfig.JavbusSite.EnableProxy {
		//set to env
		if strings.HasPrefix("http://", config.GlobalConfig.JavbusSite.ProxyUrl) || strings.HasPrefix("https://", config.GlobalConfig.JavbusSite.ProxyUrl) {
			os.Setenv("HTTP_PROXY", config.GlobalConfig.JavbusSite.ProxyUrl)
			os.Setenv("HTTPS_PROXY", config.GlobalConfig.JavbusSite.ProxyUrl)

			transport := &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					// Customize TLS settings if needed
					InsecureSkipVerify: true,
				},
			}
			// Create an HTTP client with the custom Transport
			client = &http.Client{
				Transport: transport,
				Timeout:   requestTimeoutSecond,
			}
		}

		// socks5代理
		if strings.HasPrefix("socks5://", config.GlobalConfig.JavbusSite.ProxyUrl) {
			proxyUrl := config.GlobalConfig.JavbusSite.ProxyUrl
			// Create a SOCKS5 dialer
			dialer, err := proxy.SOCKS5("tcp", proxyUrl, nil, proxy.Direct)
			if err != nil {
				fmt.Println("Error creating SOCKS5 dialer:", err)
				return nil, err
			}
			// Create a custom DialContext function
			dialContext := func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			}
			// Create a custom Transport with the SOCKS5 dialer
			transport := &http.Transport{
				DialContext: dialContext,
				TLSClientConfig: &tls.Config{
					// Customize TLS settings if needed
					InsecureSkipVerify: true,
				},
			}

			// Create an HTTP client with the custom Transport
			client = &http.Client{
				Transport: transport,
				Timeout:   requestTimeoutSecond,
			}
		}
	}
	//buid request
	requestsClient := requests.New()
	requestsClient = requestsClient.Client(client).
		Header("User-Agent", DefaultUserAgent).
		Header("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	return requestsClient, nil
}

func init() {
	client, err := NewHttpClient()
	if err != nil {
		panic(err)
	}
	HttpClient = client
}
