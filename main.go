// +build prod

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

	data, err := json.MarshalIndent(
		struct {
			Host   string        `json:"host"`
			Method string        `json:"method"`
			Header http.Header   `json:"header"`
			Body   io.ReadCloser `json:"body"`
			Query  url.Values    `json:"query,omitempty"`
		}{
			Host:   r.Host,
			Method: r.Method,
			Header: r.Header,
			Body:   r.Body,
			Query:  r.URL.Query(),
		},
		"", "\t",
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Status: Production</h1>\n\n"))
	w.Write([]byte("<h3>Engeto: Kubernetes Example Application</h3>\n\n"))

	fmt.Fprintf(w, "Request received: %s\n", data)
}
