package gfwlistpac

import "strings"

func GFWList2Pac(list GFWList) (res string) {
	data := list.ListData
	dict := generateJsDict(data)
	res = generatePac(list.DefaultProxy, dict)
	return
}

/**
 * convert bool type to string
 */
func bool2String(val bool) ( string) {
	if val {
		return "true"
	}
	return "false"
}

func generateJsDict(data GFWListData) (dic string) {
	// generate host
	dic = "var HOSTS={"
	for key, value := range data.AllowedHosts {
		if !value.HttpsEnabled && !value.HttpEnabled {
			continue
		}
		dic += "\""+key+"\":" + value.ToJavaScript()+","
	}
	dic += "};"
	// generate keywords
	dic += "var KEYWORDS=["
	for _, keyword := range data.AllowedKeywords {
		dic += "\""+strings.Replace(keyword, "\"", "\\\"", -1)+"\","
	}
	dic += "];"
	// generate exclude keywords
	dic += "var X_KEYWORDS=["
	for _, keyword := range data.ExcludedKeywords {
		dic += "\""+strings.Replace(string(keyword), "\"", "\\\"", -1)+"\","
	}
	dic += "];"
	// delete additional ","
	dic = strings.Replace(dic, ",}", "}", -1)
	dic = strings.Replace(dic, ",]", "]", -1)
	return
}

func generatePac(proxy Proxy, dict string) (pac string) {
	pac = "var FindProxyForURL = (function () {\n"
	pac += "var PROTOCOL={\"http\":"+PROTOCOL_HTTP+",\"https\":"+PROTOCOL_HTTPS+"};"
	pac += dict+"\n"
	pac += "var PROXY = '"+proxy.ToString()+"';"
	pac += `
	var DIRECT = 'DIRECT';

	function lookupDomain(host, protocol) {
		if (!host) {
			return false;
		} else if (HOSTS[host]) {
			return (HOSTS[host][PROTOCOL[protocol]]);
		} else {
			return host.indexOf('.')>0 && lookupDomain(
				host.slice(host.indexOf('.') +1),
				protocol
			);
		}
	}

	function lookupKeyword(url) {
		for (var i=0;i<X_KEYWORDS.length;i++) {
			if (!X_KEYWORDS[i] instanceof RegExp) {
				X_KEYWORDS[i] = new RegExp(X_KEYWORDS[i]);
			}
			return url.match(X_KEYWORDS[i]);
		}
		return false;
	}

	return function (url, host) {
		var doProxy = false;
		var protocol = url.substring(0, url.indexOf(':'));
		doProxy = lookupDomain(host, protocol) || lookupKeyword(url);
		return PROXY && doProxy ? PROXY : DIRECT;
	}
	})();`
	return
}
