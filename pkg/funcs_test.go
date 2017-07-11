package pkg

import "testing"

func TestSayHello(t *testing.T) {
	if SayHello() != "hello" {
		t.Fail()
	}
}