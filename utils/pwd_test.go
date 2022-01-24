package utils

import (
	"fmt"
	"testing"
)

func TestPwd(t *testing.T) {
	password, err := GenerateFromPassword("123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(CompareHashAndPassword(password, "123456"))
}
