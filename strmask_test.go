package vbcore

import "testing"

func TestStrMask(t *testing.T) {
	if StrMask("") != "" {
		t.Fail()
	}
	if StrMask("x") != "*" {
		t.Fail()
	}
	if StrMask("xx") != "x*" {
		t.Fail()
	}
	if StrMask("xxx") != "x**" {
		t.Fail()
	}
	if StrMask("xxxx") != "xx**" {
		t.Fail()
	}
}

func TestStrMaskIdx(t *testing.T) {
	if StrMaskIdx("xxxx", 0) != "****" {
		t.Fail()
	}
	if StrMaskIdx("xxxx", 1) != "x***" {
		t.Fail()
	}
	if StrMaskIdx("xxxx", 2) != "xx**" {
		t.Fail()
	}
	if StrMaskIdx("xxxx", 3) != "xxx*" {
		t.Fail()
	}
	if StrMaskIdx("xxxx", 4) != "xxxx" {
		t.Fail()
	}
}
