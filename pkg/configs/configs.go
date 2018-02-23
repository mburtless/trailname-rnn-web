package configs

import (
	"os"
)

// Map to hold cmd line args
var InstanceArgs = map[string]*string{}

// ParseFlags maps the passed cli flags to the InstanceArgs map
func ParseFlags(envVar string, parsedVal *string) {
	InstanceArgs[envVar] = parsedVal
}

// ParseEnv attempts to parse an env var whose name matches the passed
// argsKey string.  If it exists, it is saved in the InstanceArgs map with a key
// matching the env var name and the method returns true.  If it doesn't exist,
// the method returns false so the caller is aware and can handle appropriately.
func ParseEnv(argsKey string) bool {
	envKey := os.Getenv(argsKey)
	if len(envKey) > 0 {
		InstanceArgs[argsKey] = &envKey
		return true
	} else {
		return false
	}
}
