package system

import (
	"encoding/json"
	"log"
	"testing"
)

// func TestSystemDiff(t *testing.T) {
// 	SetupDiff("/home/lex", "/home")
// }

func TestGetSetupDiff(t *testing.T) {
	sud, err := GetSetupDiff(
		"mediawiki_canasta",
		"w",
		"/home/lex/mediawiki-manager/mediawiki_root/w",
		2,
	)
	if err != nil {
		log.Println(err)
	}
	jr, _ := json.MarshalIndent(sud, "  ", "  ")
	log.Print(string(jr))
}

// jr, _ := json.MarshalIndent(sud.Diff, "  ", "  ")
// log.Print(string(jr))
