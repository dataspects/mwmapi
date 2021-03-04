package mediawiki

import (
	"os"
	"testing"
)

func TestGeneralSiteInfo(t *testing.T) {
	os.Setenv("MWAPI", "https://dserver/w/api.php")
	gsi, err := GeneralSiteInfo()
	if err != nil {
		t.Errorf("Error obtaining site info")
	} else if gsi.Query.General.Base == "" {
		t.Errorf("No site info")
	}

}

func TestWfLoadExtensions(t *testing.T) {
	os.Setenv("MWROOT", "/home/lex/mediawiki-manager/mediawiki_root/w")
	wle, err := WfLoadExtensions(os.Getenv("MWROOT") + "/LocalSettings.php")
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(wle) == 0 {
		t.Errorf("No wfLoadExtensions")
	}
}
