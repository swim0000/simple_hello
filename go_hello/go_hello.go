package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := 3000

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			http.Error(w, "Failed to get hostname", http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf("Hello from %s", hostname)
		fmt.Println(message)
		w.Write([]byte(message))
	})

	fmt.Printf("Web server listening on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Failed to start the server: %v\n", err)
	}
}
