package plugin

import "testing"

func TestLoadPreset(t *testing.T) {
	a, err := LoadTextPreset("历史学家")
	if err != nil {
		t.Error("aaa：-------", err)
	}
	if a != "测试" {
		t.Error("aaa：-------", a)
	} else {
		t.Log("测试通过")
	}

}

func TestLoadMapPreset(t *testing.T) {
	a, err := LoadMapPreset("history")

	if err != nil {
		t.Error("aaa：-------", err)
	}
	for _, item := range a {
		t.Log(item.Content)
	}

}
