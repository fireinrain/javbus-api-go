package javbus

import (
	"context"
	"fmt"
	"testing"
)

func TestNewHttpClient(t *testing.T) {
	client, err := NewHttpClient()
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	var s string
	err = client.BaseURL("https://www.javbus.com").
		ToString(&s).
		Fetch(context.Background())
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	fmt.Println(s)

}

func TestNewHttpClient2(t *testing.T) {
	var s string
	err := HttpClient.BaseURL("https://www.javbus.com").
		ToString(&s).
		Fetch(context.Background())
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	fmt.Println(s)
}
