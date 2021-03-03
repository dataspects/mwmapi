package api

import (
	"net/http"

	"github.com/dataspects/mwmapi/mediawiki"
)

// Info -
func Info(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})
	resp["info"] = mediawiki.Info()
	response(w, resp, 200)
}
