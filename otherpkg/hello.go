package otherpkg

import "github.com/arekkas/accurate-test-coverage/pkg"

func HelloWorld() string {
	return pkg.SayHello() + " " + pkg.SayWorld()
}
