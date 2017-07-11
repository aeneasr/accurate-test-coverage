package otherpkg

import "testing"

func TestSayCodeCoverage(t *testing.T) {
	if SayCodeCoverage() != "code coverage" {
		t.Fail()
	}
}
