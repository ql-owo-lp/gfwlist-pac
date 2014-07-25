package gfwlistpac

import "time"

// a set of host entries
type HostEntrySet    map[string]HostEntry
type KeywordEntrySet        []string

const (
	PROTOCOL_HTTP  = "0"
	PROTOCOL_HTTPS = "1"
)

/**
 * A host entry
 */
type HostEntry struct {
	HttpEnabled        bool
	HttpsEnabled       bool

	keywordEnabled     bool
}

func (e HostEntry) ToJavaScript() (res string) {
	res += "["
	if e.HttpEnabled {
		res += "1"
	} else {
		res += "0"
	}
	res += ","
	if e.HttpsEnabled {
		res += "1"
	} else {
		res += "0"
	}
	res += "]"
	return
}


/**
 * A proxy entry
 */
type Proxy struct {
	Type               string
	Address            string
	Port               string
}

func (p Proxy) ToString() string {
	return p.Type + " " + p.Address + ":" + p.Port
}

/**
 * One GFWList
 */
type GFWList struct {
	AutoProxyTxt           string
	AutoProxyTxtMD5        string
	Output                 string
	Date                   time.Time
	ListData               GFWListData
	DefaultProxy           Proxy
}

type GFWListData struct {
	AllowedHosts           HostEntrySet
	AllowedKeywords        KeywordEntrySet
	ExcludedKeywords       KeywordEntrySet
}
