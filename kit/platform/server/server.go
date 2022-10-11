package server

import (
	"fmt"
	"net/http"
)

const (
	host = "localhost"
	port = 8000
)

var UriServer string = "localhost:8000"

func ServerRun() error {
	fmt.Println("SERVER RUN!")
	return http.ListenAndServe(UriServer, nil)
}
