package flags

var defaults flags = flags{
	IP:   "127.0.0.1",
	Port: 8080,
	Env:  "development",
	Redis: redis{
		Host:     "127.0.0.1",
		Port:     6973,
		Password: "password123",
	},
}
