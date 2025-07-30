package utils

import (
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

func LoadEnv() {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "..") // путь до корня проекта
	envPath := filepath.Join(projectRoot, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("❌ Error loading .env file from %s\n", envPath)
	}
}
