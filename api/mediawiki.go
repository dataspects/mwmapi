package api

import (
	"net/http"
	"os"

	"github.com/dataspects/mwmapi/mediawiki"
)

// MediaWikiGeneralSiteInfo -
func MediaWikiGeneralSiteInfo(w http.ResponseWriter, r *http.Request) {
	var err error
	resp := make(map[string]interface{})
	resp["data"], err = mediawiki.GeneralSiteInfo()
	resp["status"] = "MediaWiki info loaded"
	if err != nil {

	}
	response(w, resp, 200)
}

// MediaWikiWfLoadExtensions -
func MediaWikiWfLoadExtensions(w http.ResponseWriter, r *http.Request) {
	var err error
	resp := make(map[string]interface{})
	resp["info"], err = mediawiki.WfLoadExtensions(os.Getenv("MWROOT") + "/LocalSettings.php")
	if err != nil {

	}
	response(w, resp, 200)
}
