// main.go
package main

import (
	"flights/pkg/server"
	"fmt"
	"net/http"
	"os"
)

func main() {
	serverStart()
}

func serverStart() {
	var err error

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	mux := http.NewServeMux()

	calculateHandler := server.MiddlewarePost(http.HandlerFunc(server.GeneratesSubRoutesOfRoute))
	mux.Handle("/calculate", calculateHandler)

	fmt.Printf("Server started at http://localhost%v\n", port)
	if err = http.ListenAndServe(port, mux); err != nil {
		panic(fmt.Errorf("main.http.ListenAndServe().error: %v", err))
	}
}
