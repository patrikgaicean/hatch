package config

type Config struct {
	Port        int
	Env         string
	RegistryURL string
	// Other configuration properties...
}

// New creates a new instance of Config and returns a pointer to it.
func New() *Config {
	return &Config{
		Port:        8080,
		Env:         "develop",
		RegistryURL: "someurl",
		// Initialize properties as needed, e.g., Port: 8080
	}
}
