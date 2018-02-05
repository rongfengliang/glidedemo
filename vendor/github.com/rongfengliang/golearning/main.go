package main

import (
	"fmt"

	"github.com/rongfengliang/golearning/server"
)

const appversion = "1.0"

func main() {
	// var user = users.User{Name: "dalofffffffffng", Age: 44}
	// fmt.Printf("dalong" + "世界")
	// fmt.Printf(user.Name)
	// resp, err := http.Get("http://www.baidu.com")
	// if err != nil {
	// 	fmt.Printf("there is some wrong for get data")
	// }
	// b, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	fmt.Printf("wrong convert data")
	// 	os.Exit(1)
	// }
	// fmt.Printf("%s", b)
	fmt.Print(appversion)
	server := server.Server{}
	server.HttpServer()
}
