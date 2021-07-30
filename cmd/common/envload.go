package common

import (
	"flag"

	"github.com/joho/godotenv"
)

var (
	EnvFile = flag.String("e", "", "set the env file to read")
)

func init() {
	flag.Parse()
}

func LoadEnv() error {
	if err := godotenv.Load(*EnvFile); err != nil {
		return err
	}
	return nil
}
