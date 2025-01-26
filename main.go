package main

import (
	"fmt"
	handler "heck2api/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)

	port := "3000"
	fmt.Printf("Server Start At http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Server Failed: %v\n", err)
	}
}

