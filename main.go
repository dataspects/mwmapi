package main

import (
	"log"

	"github.com/dataspects/mwmapi/http"
)

func main() {
	err := http.Serve(3000, []string{"https://dserver/ui/", "http://dserver:8000"})
	if err != nil {
		log.Println(err)
	}
}
