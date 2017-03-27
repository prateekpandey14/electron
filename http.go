package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func serveHTTP(port int, host string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestHandler(w, r)
	})

	addr := fmt.Sprintf("%v:%d", host, port)
	server := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	log.Println(err.Error())

}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `M-Api Server is Running!`)
}

func main() {

	go serveHTTP(5656, "127.0.0.1")

	select {}
}
