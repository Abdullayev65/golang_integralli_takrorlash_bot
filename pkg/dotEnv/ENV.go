package dotEnv

import (
	"fmt"
	"github.com/joho/godotenv"
)

var EnvMap map[string]string

func init() {
	envMap, err := godotenv.Read("my_env")
	if err != nil {
		fmt.Println(".env NOT FOUND")
		panic(err)
	}
	EnvMap = envMap
}
