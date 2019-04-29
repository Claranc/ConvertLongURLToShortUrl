package controller

import "testing"

func TestConvertTenToOtherjinzhi(t *testing.T) {
	var input = []int{5,9,15,3}
	var output = []string{"101","1001","1111","11" }
	for i := 0; i < len(input); i++ {
		if output[i] != convertTenToOtherJinzhi(input[i], 2) {
			t.Fatal("wrong")
		}
	}
}


func TestCheckValidOfLongUrl(t *testing.T) {
	var s = []string{"http://www.baidu.com", "fdasfdasfads", "https://$^&^&*.fdafadsf.fdasf", "http://www.dfsafdasfads.com"}
	var res = []bool{true, false, false,true}
	for i := 0; i < len(s); i++ {
		if CheckValidOfLongUrl(s[i]) != res[i] {
			t.Fail()
			t.Log(s[i])
		}
	}
}