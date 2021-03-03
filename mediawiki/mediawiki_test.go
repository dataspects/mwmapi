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
		t.Logf("")
	}

}
