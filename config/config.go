package config

import (
	"os"
)

//Config is the struct of configurations
type Config struct {
	Env        string
	StatusCode int
}

//configuration variables
var (
	Port             string
	HttpReadTimeout  string
	HttpWriteTimeout string
	GroupRequest     string
)

func init() {
	Port = os.Getenv("PORT")
	HttpReadTimeout = os.Getenv("HTTP_READ_TIMEOUT")
	HttpWriteTimeout = os.Getenv("HTTP_WRITE_TIMEOUT")
	GroupRequest = os.Getenv("GROUPREQUEST")
}
