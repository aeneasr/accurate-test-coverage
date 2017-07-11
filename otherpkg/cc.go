package otherpkg

import "github.com/arekkas/accurate-test-coverage/pkg"

func CodeCoverage() string {
	return pkg.SayCode() + " " + pkg.SayCoverage()
}
