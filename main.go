package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	host = "0.0.0.0"
	port = "8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)

	if val, exists := os.LookupEnv("SERVER_HOST"); exists {
		host = val
	}
	if val, exists := os.LookupEnv("SERVER_PORT"); exists {
		port = val
	}

	addr := host + ":" + port
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	fmt.Printf("Serving Echo Server on: http://%s\n", addr)

	done := make(chan bool)
	go func() {
		// FIXME: ServeTLS timeouts on handshake
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
		fmt.Print("Stopped serving livesport")
		done <- true
	}()
	<-done
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Engeto: Kubernetes Example Application\n\n"))

	fmt.Fprintf(w, "Request received: %#v", r)
}
