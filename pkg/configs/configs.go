package configs

/*import (
	"os"
)*/

// Map to hold cmd line args
var InstanceArgs = map[string]*string{}

func ParseFlags(apiHost *string) {
	InstanceArgs["ApiHost"] = apiHost
}

/*func ParseEnv(envKey string) {
	
}*/
