package mediawiki

import (
	"testing"
)

// Info -
func TestGeneralSiteInfo(t *testing.T) {
	gsi, err := GeneralSiteInfo()
	if err != nil {
		t.Errorf(err.Error())
	}
	if gsi.Query.General.Base != "" {
		t.Logf("OK")
	} else {
		t.Errorf("Failed")
	}

}

func TestWfLoadExtensions(t *testing.T) {
	wle, err := WfLoadExtensions("/home/lex/mediawiki-manager/mediawiki_root/w/LocalSettings.php")
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(wle) > 0 {
		t.Logf("OK")
	} else {
		t.Errorf("Failed")
	}
}
