package main

import (
	"fmt"
	"net/http"
	"os"

	"clstr.co/gbdfail"
)

func main() {
	mux := http.NewServeMux()
	gbdfail.Routes(mux)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Fprintf(os.Stderr, "listen and serve: %v\n", err)
		os.Exit(1)
	}
}
