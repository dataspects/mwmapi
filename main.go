package main

import (
	"log"

	"github.com/dataspects/mwmapi/http"
)

func main() {
	err := http.Serve(3002, []string{"https://dserver/ui/", "http://dserver:8001"})
	if err != nil {
		log.Println(err)
	}
}
