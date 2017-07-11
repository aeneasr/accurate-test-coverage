package otherpkg

import "github.com/arekkas/accurate-test-coverage/pkg"

func SayCodeCoverage() string {
	return pkg.SayCode() + " " + pkg.SayCoverage()
}
