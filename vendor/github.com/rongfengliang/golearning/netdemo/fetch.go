package netdemo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetContentFromNet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("there is some wrong for get data")
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("wrong convert data")
		os.Exit(1)
	}
	str2 := string(b[:])
	return str2
}
