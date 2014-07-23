package gfwlistpac

import "time"

// a set of host entries
type HostEntrySet 	map[string]HostEntry
type KeywordEntrySet		[]KeywordEntry

/**
 * A host entry
 */
type HostEntry struct {
	Protocol		string
	ProxySelected	Proxy

	httpEnabled      bool
	httpsEnabled     bool
	httpDisabled     bool
	httpsDisabled    bool
	keywordEnabled   bool
}

type KeywordEntry string

/**
 * A proxy entry
 */
type Proxy struct {
	Type			string
	Address			string
	Port			string
}

/**
 * One GFWList
 */
type GFWList struct {
	AutoProxyTxt		string
	AutoProxyTxtMD5		string
	PacTxt				string
	Date				time.Time
	Data				GFWListData
}

type GFWListData struct {
	AllowedHosts		HostEntrySet
	AllowedKeywords		KeywordEntrySet
	ExcludedKeywords	KeywordEntrySet
}
