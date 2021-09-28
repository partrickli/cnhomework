package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	ver := os.Getenv("VERSION")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Log client IP, HTTP response code
		ip := r.RemoteAddr

		fmt.Println("Client IP Address: ", ip) //

		for name, values := range r.Header {
			w.Header().Set(name, strings.Join(values, ","))
			for _, value := range values {
				fmt.Println(name, value)
			}
		}
		fmt.Println(ver)
		w.Header().Set("VERSION", ver)

		w.WriteHeader(http.StatusNoContent)
	})

	http.HandleFunc("/localhost/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
