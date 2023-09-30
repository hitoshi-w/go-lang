package main

import (
	"fmt"
	"net/http"
	"os"
)


func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %s", r.URL.Path[1:])
	})

	err := http.ListenAndServe(":3000", handler)

	if err != nil {
		fmt.Printf("failed to server: %v", err)
		os.Exit(1)
	}
}