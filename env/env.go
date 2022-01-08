package env

import (
	"os"
	"strconv"
	"strings"
)

var defaults map[string]string

func init() {
	defaults = map[string]string{
		"ipv4_ip":            "10.0.0.1",
		"ipv4_n":             "16",
		"db_user":            "root",
		"db_pwd":             "123456",
		"db_host":            "127.0.0.1:3305",
		"db_name":            "auth",
		"db_scheme":          "mysql",
		"db_show_sql":        "false",
		"dhcp_default_group": "dhcp0",
		"dhcp_default_num":   "60000",
		"dhcp_default_ip":    "10.0.0.1",
	}
}

func getDefault(key string) string {
	return defaults[key]
}

func Get(key string) string {
	env := strings.TrimSpace(os.Getenv(strings.ToUpper(key)))
	if env != "" {
		return env
	}
	return getDefault(key)
}
func GetInt(key string) int {
	v, _ := strconv.Atoi(Get(key))
	return v
}
func GetInt64(key string) int64 {
	return int64(GetInt(key))
}

func GetBool(key string) bool {
	return "true" == strings.ToLower(Get(key))
}
