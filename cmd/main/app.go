package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, par httprouter.Params) {
	name := par.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello is %s", name)))
}

func main() {
	router := httprouter.New()
	router.GET("/:name", IndexHandler)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))

}
