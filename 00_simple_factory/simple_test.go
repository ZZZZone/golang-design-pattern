package simplefactory

import (
	"testing"
)

//TestType1 test get hiapi with factory
func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		t.Fatal("Type2 test fail")
	}
}

func TestType3(t *testing.T) {
	api := NewAPI(3)
	s := api.Say("张三")
	if s != "你好, 张三" {
		t.Fatal("Type2 test fail")
	}
}

func TestOther(t *testing.T) {
	api := NewAPI(4)
	if api != nil {
		t.Fatal("api is not nil")
	}
}
