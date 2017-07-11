package pkg

import "testing"

func TestSayCode(t *testing.T) {
	if SayCode() != "code" {
		t.Fail()
	}
}