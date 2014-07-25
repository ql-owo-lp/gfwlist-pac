package gfwlistpac

import (
	"regexp"
	"encoding/base64"
	"net/http"
	"io/ioutil"
	"appengine/urlfetch"
	"appengine"
)

// official URL of GFWList
const GFWListURL = "http://autoproxy-gfwlist.googlecode.com/svn/trunk/gfwlist.txt"

func FetchGFWListDesktop() string {
	resp, err := http.Get(GFWListURL)
	return fetchGFWList(resp, err)
}

func FetchGFWListGAE(w http.ResponseWriter, r *http.Request) string {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(GFWListURL)
	return fetchGFWList(resp, err)
}

/**
 * Fetch GFW list from official site
 */
func fetchGFWList(resp *http.Response, err error) (list string) {
	whitespaceRegex := regexp.MustCompile("[\t \r\n]")
	listRegex := regexp.MustCompile("![^\n]+\n")
	emptyLn := regexp.MustCompile("\n\n")

	if err != nil {
		resp.Body.Close()
		return ""
	}
	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return ""
	}

	str := whitespaceRegex.ReplaceAllString(string(content), "")
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}
	list = string(data)
	list = listRegex.ReplaceAllString(list, "")
	list = emptyLn.ReplaceAllString(list, "\n")
	return
}
