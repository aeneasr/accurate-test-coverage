package otherpkg

import "testing"

func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "hello world" {
		t.Fail()
	}
}