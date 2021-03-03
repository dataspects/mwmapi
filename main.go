package main

import (
	"log"

	"github.com/dataspects/mwmapi/http"
)

func main() {
	err := http.Serve("8080", []string{"localhost"})
	if err != nil {
		log.Println(err)
	}
}
