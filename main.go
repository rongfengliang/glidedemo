package main

import (
	"fmt"

	"github.com/Masterminds/cookoo"
	"github.com/rongfengliang/golearning/server"
)

func init() {
	fmt.Print("init main")
}

var directives = []string{
	"example",
	"whoami",
}

func main() {
	// Build a new Cookoo app.
	p := new(server.Server)
	p.HttpServer()
	registry, router, context := cookoo.Cookoo()

	// Fill the registry.
	registry.AddRoutes(
		cookoo.Route{
			Name: "TEST",
			Help: "A test route",
			Does: cookoo.Tasks{
				cookoo.Cmd{
					Name: "hi",
					Fn:   HelloWorld,
				},
			},
		},
	)

	// Execute the route.
	router.HandleRequest("TEST", context, false)

}
func HelloWorld(cxt cookoo.Context, params *cookoo.Params) (interface{}, cookoo.Interrupt) {
	fmt.Println("Hello World")
	return true, nil
}
