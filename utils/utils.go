package utils

import (
	"os"
)

func GetValue(key string) string {
	return os.Getenv(key)
}
