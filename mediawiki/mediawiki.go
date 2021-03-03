package mediawiki

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// GSI -
type GSI struct {
	Query Query `json:"query"`
}

// Query -
type Query struct {
	General    General     `json:"general"`
	Extensions []Extension `json:"extensions"`
	Skins      []Skin      `json:"skins"`
}

// General -
type General struct {
	Base          string `json:"base"`
	Dbtype        string `json:"dbtype"`
	Dbversion     string `json:"dbversion"`
	Generator     string `json:"generator"`
	Mainpage      string `json:"mainpage"`
	Maxuploadsize int    `json:"maxuploadsize"`
	Phpversion    string `json:"phpversion"`
	Sitename      string `json:"sitename"`
	Time          string `json:"time"`
	Timeoffset    int    `json:"timeoffset"`
	Timezone      string `json:"timezone"`
}

// Extension -
type Extension struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

// Skin -
type Skin struct {
	Code string `json:"code"`
}

// GeneralSiteInfo -
func GeneralSiteInfo() (GSI, error) {
	gsi := GSI{}
	props := []string{
		"general",
		"extensions",
		"skins",
	}
	res := mwapicall("https://dserver/w/api.php?format=json&action=query&meta=siteinfo&siprop=" + strings.Join(props, "|"))
	err := json.NewDecoder(res.Body).Decode(&gsi)
	if err != nil {

	}
	defer res.Body.Close()
	log.Println(gsi)

	return gsi, nil
}

func mwapicall(url string) *http.Response {
	client := client()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

	}
	res, err := client.Do(req)
	if err != nil {

	}
	return res
}

func client() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return client
}
