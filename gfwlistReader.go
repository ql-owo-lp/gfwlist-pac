package gfwlistpac

import (
	"regexp"
	"strings"
)

/**
 * convert GFW to raw data structures
 */
func ReadGFWList(listContent string) (data GFWListData) {
	const DOMAIN_REGEX_WITHOUT_EOL = "^(?:(?:(?:@@)?\\|?(?:(?:https?:\\/\\/)|\\|))?\\.?((?:[\\-A-Za-z0-9]+\\.)+[A-Za-z]+)\\/?|((?:\\d{1,3}\\.){3}\\d{1,3}))"

	data.AllowedHosts = make(HostEntrySet)
	urls := strings.Split(listContent, "\n")
	domainRegex := regexp.MustCompile(DOMAIN_REGEX_WITHOUT_EOL)
	pureDomainRegex := regexp.MustCompile(DOMAIN_REGEX_WITHOUT_EOL + "$")
	regexpRegex := regexp.MustCompile("^\\/(.+)\\/$")
	// notice here the first line will be [AutoProxy version]
	for i := 1; i < len(urls); i++ {
		curUrl := urls[i]
		if curUrl == "" {
			continue
		}
		host := ""
		hostMatches := domainRegex.FindStringSubmatch(curUrl)
		if len(hostMatches) == 3 {
			host = hostMatches[1]
			if host == "" {    // in case it's IP
				host = hostMatches[2]
			}
		}

		//				fmt.Println(len(hostMatches), domainRegex.FindStringSubmatch(curUrl))
		//				fmt.Println(curUrl, "->", host)
		if host == "" {    // no match, this could be a keyword
			if matches := regexpRegex.FindStringSubmatch(curUrl); len(matches) == 2 {
				data.AllowedKeywords = append(data.AllowedKeywords, matches[1])
				//				fmt.Println(matches[1])
			} else {
				// convert to regex format
				tmpStr := strings.Replace(curUrl, ".", "\\.", -1)
				if strings.HasPrefix(curUrl, "@@") {
					tmpStr = strings.Replace(tmpStr, "@@", "", 1)
				}
				tmpStr = strings.Replace(tmpStr, "?", "\\?", -1)
				tmpStr = strings.Replace(tmpStr, "^", "\\^", -1)
				tmpStr = strings.Replace(tmpStr, "$", "\\$", -1)
				tmpStr = strings.Replace(tmpStr, "*", ".*", -1)
				if strings.HasPrefix(tmpStr, "|http") {
					tmpStr = strings.Replace(tmpStr, "|http", "^http", 1)
				} else if strings.HasPrefix(tmpStr, "||") {
					tmpStr = strings.Replace(tmpStr, "||", "", 1)
				} else if strings.HasPrefix(tmpStr, "|") {
					tmpStr = strings.Replace(tmpStr, "|", "", 1)
				} else {
					tmpStr = "^http://.*"+tmpStr
				}
				tmpStr = strings.Replace(tmpStr, "|", "\\|", -1)
				tmpStr = strings.Replace(tmpStr, "/", "\\/", -1)
				if strings.HasPrefix(curUrl, "@@") {
					// exclude list
					data.ExcludedKeywords = append(data.ExcludedKeywords, tmpStr)
				} else {
					data.AllowedKeywords = append(data.AllowedKeywords, tmpStr)
				}
				//				fmt.Println(curUrl, "->", tmpStr)
			}
			continue    // do not create host entry here
		}

		var entry HostEntry
		// get existing entry if any
		if val, ok := data.AllowedHosts[host]; ok {
			entry = val
		} else {
			// set default value here
			entry = HostEntry{
				HttpEnabled : false,
				HttpsEnabled : false,
				keywordEnabled : false,
			}
		}

		if pureDomainRegex.MatchString(curUrl) {
			// ||domain , means accept both http and https protocol
			if strings.HasPrefix(curUrl, "||") {
				entry.HttpEnabled = true
				entry.HttpsEnabled = true
			} else if strings.HasPrefix(curUrl, "|http://") {
				entry.HttpEnabled = true
			} else if strings.HasPrefix(curUrl, "|https://") {
				entry.HttpsEnabled = true
			} else if strings.HasPrefix(curUrl, "@@||") {
				// overwrite existing
				entry.HttpEnabled = false
				entry.HttpsEnabled = false
			}else if strings.HasPrefix(curUrl, "@@|http://") {
				entry.HttpEnabled = false
			} else if strings.HasPrefix(curUrl, "@@|https://") {
				entry.HttpsEnabled = false
			} else {    // plain domain / keyword
				entry.HttpEnabled = true
				entry.keywordEnabled = true
				// put this entry into keyword list as well
				//			keywordList = append(keywordList, curUrl)
			}
		}
		//		fmt.Println(curUrl,"->",host,":",entry)
		data.AllowedHosts[host] = entry
	}

	return
}
