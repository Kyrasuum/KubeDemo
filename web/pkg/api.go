package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var (
	api_host = condAssignStrEnv("API_HOST", "http://127.0.0.1:8081")
)

func condAssignStrEnv(env string, def string) string {
	val := os.Getenv(env)
	if val == "" {
		fmt.Printf("Unset: %s, using: %s\n", env, def)
		val = def
	}
	return val
}

func condAssignIntEnv(env string, def int) int {
	varstr := os.Getenv(env)
	if varstr == "" {
		fmt.Printf("Unset: %s, using: %d\n", env, def)
		return def
	}
	return validInt(varstr, env, def)
}

func validInt(variable string, name string, def int) int {
	value, err := strconv.Atoi(variable)
	if err != nil {
		fmt.Printf("Invalid integer value for %s, setting to %d\n", name, def)
		value = def
	}

	return value
}

func GetData(path string) ([]byte, error) {
	resp, err := http.Get(api_host + path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("open %s: file does not exist", path)
	}

	return ioutil.ReadAll(resp.Body)
}
