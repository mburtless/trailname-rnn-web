package configs

import (
	"os"
)

// Struct to store values for configuration details for each flag or env var
// used to configure this app
type ConfigSet struct {
	DefVal	string
	EnvVar	string
	FlagVar string
	Help	string
	ParsedVal	string
}

// Map that defines each flag or env var used to configure this app
var ConfigVars = map[string]*ConfigSet{
	"apiHost":		{"localhost", "APIHOST", "apihost", "IP or hostname of the trailname-rnn API", ""},
	"port":			{"8000", "PORT", "port", "Port to listen on", ""},
	"environment":	{"development", "ENVIRONMENT", "environemnt", "Deployment environment", ""},
}
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
