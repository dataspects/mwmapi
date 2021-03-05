package mediawiki

import (
	"log"
	"os"
	"testing"
)

func TestIsInSafeMode(t *testing.T) {
	iism, _ := IsInSafeMode(os.Getenv("MWCONTAINER"))
	log.Println(iism)
}
func TestGeneralSiteInfo(t *testing.T) {
	gsi, err := GeneralSiteInfo()
	if err != nil {
		t.Errorf("Error obtaining site info")
	} else if gsi.Query.General.Base == "" {
		t.Errorf("No site info")
	}

}

func TestWfLoadExtensions(t *testing.T) {
	wle, err := WfLoadExtensions(os.Getenv("MWROOT") + "/LocalSettings.php")
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(wle) == 0 {
		t.Errorf("No wfLoadExtensions")
	}
}
