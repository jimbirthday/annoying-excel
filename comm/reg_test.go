package comm

import (
	"regexp"
	"testing"
)

func TestRegUAP(t *testing.T) {
	compile := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{4,15}$`)
	println("user :", compile.MatchString("jim_bir"))

	compiles := regexp.MustCompile(`^[a-zA-Z]\w{5,17}$`)
	println("pwd :", compiles.MatchString("A1234_567"))
}
