package gfwlistpac

import "time"

// a set of host entries
type HostEntrySet    map[string]HostEntry
type KeywordEntrySet        []string

/**
 * A host entry
 */
type HostEntry struct {
	Protocol         string
	ProxySelected    Proxy

	httpEnabled      bool
	httpsEnabled     bool
	httpDisabled     bool
	httpsDisabled    bool
	keywordEnabled   bool
}

/**
 * A proxy entry
 */
type Proxy struct {
	Type               string
	Address            string
	Port               uint16
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
