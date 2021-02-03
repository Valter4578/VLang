package object

import "testing"

func TestStringHashKey(t *testing.T) {
	example1 := &String{Value: "Hello World"}
	example2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if example1.HashKey() != example2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if example1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
