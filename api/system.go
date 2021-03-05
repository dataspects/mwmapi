package api

import (
	"net/http"

	"github.com/dataspects/mwmapi/system"
)

// SystemGetSetupDiff -
func SystemGetSetupDiff(w http.ResponseWriter, r *http.Request) {
	var err error
	resp := make(map[string]interface{})
	resp["data"], err = system.GetSetupDiff("mediawiki_canasta",
		"w",
		"/home/lex/mediawiki-manager/mediawiki_root/w",
		2)
	resp["status"] = "GetSetupDiff loaded"
	if err != nil {

	}
	response(w, resp, 200)
}
