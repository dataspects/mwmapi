package mediawiki

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
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
	log.Println("Requesting GeneralSiteInfo...")
	gsi := GSI{}
	props := []string{
		"general",
		"extensions",
		"skins",
	}
	res := mwapicall(os.Getenv("MWAPI") + "?format=json&action=query&meta=siteinfo&siprop=" + strings.Join(props, "|"))
	err := json.NewDecoder(res.Body).Decode(&gsi)
	if err != nil {

	}
	defer res.Body.Close()
	return gsi, nil
}

// WfLoadExtensions -
func WfLoadExtensions(lsURL string) ([]string, error) {
	log.Println("Requesting WfLoadExtensions...")
	wle := []string{}
	data, err := ioutil.ReadFile(lsURL)
	if err != nil {
		return wle, err
	}
	re := regexp.MustCompile("#?wfLoadExtension.*;")
	matches := re.FindAllString(string(data), -1)

	for _, match := range matches {
		wle = append(wle, match)
	}
	return wle, nil
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
