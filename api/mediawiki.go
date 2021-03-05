package api

import (
	"net/http"
	"os"

	"github.com/dataspects/mwmapi/mediawiki"
	"github.com/dataspects/mwmapi/system"
)

// MediaWikiIsInSafeMode -
func MediaWikiIsInSafeMode(w http.ResponseWriter, r *http.Request) {
	var err error
	resp := make(map[string]interface{})
	iism, err := mediawiki.IsInSafeMode(os.Getenv("MWCONTAINER"))
	if iism {
		sud, err := system.GetSetupDiff(
			os.Getenv("MWCONTAINER"),
			"w",
			os.Getenv("INSTANCEMWROOT"),
			2,
		)
		if err != nil {

		}
		resp["setupDiff"] = sud
	}
	resp["isInSafeMode"] = iism
	resp["status"] = "MediaWikiIsInSafeMode info loaded"
	if err != nil {

	}
	response(w, resp, 200)
}

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
