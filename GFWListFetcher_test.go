package gfwlistpac

import (
	"testing"
	"fmt"
)

func TestFetchGFWList(t *testing.T) {
	list := FetchGFWListDesktop()
	fmt.Println(list)
}
