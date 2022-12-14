package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type configEnv struct {
}

type Config interface {
	Get(string) any
	GetString(string) string
	GetInt(string) int
}

func LoadEnv() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &configEnv{}
}

func (c *configEnv) Get(env string) any {
	data := os.Getenv(env)
	return data
}

func (c *configEnv) GetString(env string) string {
	data := c.Get(env).(string)
	return data
}

func (c *configEnv) GetInt(env string) int {
	data := c.Get(env).(int)
	return data
}
