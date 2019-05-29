package main

import "testing"

func TestADD(t *testing.T) {
	r := add(3, 6)
	if r != 9 {
		t.Fatalf("cuowu:%d",r)
	}
	t.Logf("succ")
}


func TestSub(t *testing.T) {

	r := sub(2, 4)
	if r != -2 {
		t.Fatalf("sub(2, 4) error, expect:%d, actual:%d", 6, r)
	}
	t.Logf("test sub succ")

}
