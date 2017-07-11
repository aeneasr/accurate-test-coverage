package otherpkg

import "testing"

func TestCodeCoverage(t *testing.T) {
	if CodeCoverage() != "code coverage" {
		t.Fail()
	}
}
