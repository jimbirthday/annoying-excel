package utils

import (
	"fmt"
	ginpro "github.com/gin-pro/gin-pro-base"
	"os"
)

func GenerateNick() string {
	return fmt.Sprintf("%d_%s", os.Getpid(), ginpro.RandomString(5))
}
