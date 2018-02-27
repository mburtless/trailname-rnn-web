package configs

import (
	"os"
	"log"
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
	"apiHost":		{"localhost:6788", "APIHOST", "apihost", "IP or hostname and port of the trailname-rnn API", ""},
	"port":			{"8000", "PORT", "port", "Port to listen on", ""},
	"environment":	{"development", "ENVIRONMENT", "environemnt", "Deployment environment", ""},
}
// Map to hold cmd line args
var InstanceArgs = map[string]*string{}

func ParseConfigVars() {
	// If ParseEnv indicates APIHOST env doesn't exist, store user value or default passed via flag
	for key, val := range ConfigVars {
		if ParseEnv(key, val.EnvVar) {
			log.Printf("%s env found, ignoring flag value and using %s instead", val.EnvVar, val.ParsedVal)
		} else {
			// Send to configs for parsing into InstanceArgs
			log.Printf("%s env not found, using flag value of %s instead", val.EnvVar, val.ParsedVal)
			//ParseFlags(val.EnvVar, &val.ParsedVal)
		}
	}
}

// ParseFlags maps the passed cli flags to the InstanceArgs map
func ParseFlags(envVar string, parsedVal *string) {
	InstanceArgs[envVar] = parsedVal
}

// ParseEnv attempts to parse an env var whose name matches the passed
// argsKey string.  If it exists, it is saved as the ParsedVal in the ConfigVars map
// and the method returns true.  If it doesn't exist, the method returns false so the
// caller is aware and can handle appropriately.
func ParseEnv(configVar string, envVar string) bool {
	envVal := os.Getenv(envVar)
	if len(envVal) > 0 {
		//InstanceArgs[argsKey] = &envKey
		ConfigVars[configVar].ParsedVal = envVal
		return true
	} else {
		return false
	}
}
