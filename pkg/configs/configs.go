package configs

// Struct to hold cmd line args
type Args struct {
	ApiHost		*string
}

// cmd line args for this instance
var InstanceArgs Args

func ParseFlags(apiHost *string) {
	InstanceArgs.ApiHost = apiHost
}
