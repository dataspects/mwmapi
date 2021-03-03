package api

import (
	"net/http"

	"github.com/dataspects/mwmapi/mediawiki"
)

// MediaWikiGeneralSiteInfo -
func MediaWikiGeneralSiteInfo(w http.ResponseWriter, r *http.Request) {
	var err error
	resp := make(map[string]interface{})
	resp["info"], err = mediawiki.GeneralSiteInfo()
	if err != nil {

	}
	response(w, resp, 200)
}
