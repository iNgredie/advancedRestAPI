package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)


func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))

}

func main()  {
	router := httprouter.New()
	router.GET("/:name", IndexHandler)

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil{
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}