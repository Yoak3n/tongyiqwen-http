package util

import "testing"

func TestOutLine(t *testing.T) {
	//t.Error("Not implemented")
	result := OutLine("中国封建社会通常\n被认为起\n始于公\n元前8世纪的周朝，终")
	if result == nil {
		t.Error("Not implemented")
	} else {
		t.Log(result)
	}

}
