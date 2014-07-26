package gfwlistpac

import (
	"net/http"
	"fmt"

)

func init() {
	http.HandleFunc("/pac", genProxy)
}

// [END func_root]

func genProxy(w http.ResponseWriter, r *http.Request) {
	gfwlistTxt := FetchGFWListGAE(w, r)
	gfwlistDat := ReadGFWList(gfwlistTxt)

	defaultProxy := Proxy {
		Type : "SOCKS5",
		Address : "127.0.0.1",
		Port : "8088",
	}

	gfwlist := GFWList {
		DefaultProxy : defaultProxy,
		AutoProxyTxt : gfwlistTxt,
		AutoProxyTxtMD5 : "",
		ListData : gfwlistDat,
	}
	gfwlist.Output = GFWList2Pac(gfwlist)

	w.Header().Set("Content-Type", "application/x-ns-proxy-autoconfig")

	fmt.Fprintf(w, "%s", gfwlist.Output)
}
