package config

import "os"

func GetApiPort() {
	os.Getenv("API_PORT")
}
