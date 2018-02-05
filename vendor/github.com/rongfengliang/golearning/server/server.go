package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rongfengliang/golearning/netdemo"
)

type Server struct {
}

func (s Server) Indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", netdemo.GetContentFromNet("http://www.cnblogs.com/rongfengliang"))
}
func (s Server) HttpServer() {
	http.HandleFunc("/", s.Indexhandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
