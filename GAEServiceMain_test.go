package gfwlistpac

import "testing"

func TestGFWList2Pac(t *testing.T) {
	gfwlistTxt := FetchGFWList()
	gfwlistDat := ReadGFWList(gfwlistTxt)

	defaultProxy := Proxy {
		Type : "SOCKS5",
		Address : "127.0.0.1",
		Port : 8088,
	}

	gfwlist := GFWList{
		DefaultProxy : defaultProxy,
		AutoProxyTxt : gfwlistTxt,
		AutoProxyTxtMD5 : "",
		ListData : gfwlistDat,
	}
	gfwlist.Output = GFWList2Pac(gfwlist)
}
