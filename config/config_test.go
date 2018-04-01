package config

import "testing"

func TestGetString(t *testing.T) {
	v, err := GetString("appname")
	if err != nil {
		t.Errorf("GetString error:%v\n", err)
	}
	if v == "MyDoc" {
		t.Log("GetString OK")
	} else {
		t.Fail()
	}
}

func TestGetStrings(t *testing.T) {
	v, err := GetStrings("saveDir")
	testString := []string{"abc", "def"}
	if err != nil {
		t.Errorf("GetString error:%v\n", err)
	}
	for i := range v {
		if testString[i] != v[i] {
			t.Fail()
		}
	}
}

func TestGetInt(t *testing.T) {
	v, err := GetInt("port")
	if err != nil {
		t.Errorf("TestGetInt error:%v\n", err)
	}
	if v == 8080 {
		t.Log("GetString OK")
	} else {
		t.Fail()
	}
}
