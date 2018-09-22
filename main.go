package main

import (
	"fmt"

	"github.com/mpiercy827/apod-api-go/server"
)

func main() {
	srv := server.New()

	err := srv.ListenAndServe()

	if err != nil {
		fmt.Sprintf("Error: %s", err)
	}
}
