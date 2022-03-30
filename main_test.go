package main

import "testing"

func TestGeneral(t *testing.T) {
	if true {
		t.Fail()
		t.Log("test should pass")
	}
}
