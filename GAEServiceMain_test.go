package gfwlistpac

import "testing"

func TestGFWList2Pac(t *testing.T) {
	gfwlistTxt := FetchGFWList()
	gfwlistDat := ReadGFWList(gfwlistTxt)
	gfwlist := GFWList{
		DefaultProxy : "",
		AutoProxyTxt : gfwlistTxt,
		AutoProxyTxtMD5 : "",
		ListData : gfwlistDat,
	}
	gfwlist.Output = GFWList2Pac(gfwlist)
}
