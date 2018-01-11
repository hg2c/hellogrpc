package lib

import (
	"os"
	"strings"
)

func getenv(key string) string {
	k := strings.ToUpper(key)
	k = strings.Join(strings.Split(k, "."), "_")
	v := os.Getenv(k)
	return v
}

func GetConfig(key string, default_value ...string) string {
	var v string

	v = getenv(key)
	if v != "" {
		return v
	}

	if len(default_value) > 0 {
		v = default_value[0]
	}
	return v
}
