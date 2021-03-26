package config

// DatabaseConfiguration contains our SQL database configuration
type DatabaseConfiguration struct {
	// ...
}

// GRPCServerConfiguration contains our GRPC configuration
type GRPCServerConfiguration struct {
	Port int
	// ...
}

// WebServerConfiguration contains the configuration related to the
// web server exposing our REST API.
type WebServerConfiguration struct {
	Port int
	// ...
}

// GetDatabaseConfiguration returns our web server configuration
func GetDatabaseConfiguration() DatabaseConfiguration {
	return DatabaseConfiguration{}
}

// GetGRPCServerConfiguration returns our GRPC server configuration
func GetGRPCServerConfiguration() GRPCServerConfiguration {
	return GRPCServerConfiguration{}
}

// GetWebServerConfiguration returns our web server configuration
func GetWebServerConfiguration() WebServerConfiguration {
	// Here we would populate our configuration from the config file
	// we read in cmd/main
	return WebServerConfiguration{}
}
