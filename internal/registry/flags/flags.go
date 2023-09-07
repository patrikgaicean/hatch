package flags

import (
	"flag"
	"fmt"
	"math"
)

type flags struct {
	IP      string
	Port    uint16
	Env     string
	Cleanup int64
	Redis   redis
}

type redis struct {
	Host     string
	Port     uint16
	Password string
}

var defaults flags = flags{
	IP:      "127.0.0.1",
	Port:    8080,
	Env:     "dev",
	Cleanup: 15,
	Redis: redis{
		Host:     "127.0.0.1",
		Port:     6973,
		Password: "password123",
	},
}

func ParseFlags() flags {
	var (
		f          flags
		serverPort uint
		redisPort  uint
	)

	flag.StringVar(
		&f.IP,
		"ip",
		defaults.IP,
		"Registry ip",
	)
	flag.UintVar(
		&serverPort,
		"port",
		uint(defaults.Port),
		"Registry port",
	)
	flag.StringVar(
		&f.Env,
		"env",
		defaults.Env,
		"Environment (dev|stg|prod)",
	)
	flag.Int64Var(
		&f.Cleanup,
		"cleanup",
		defaults.Cleanup,
		"Cleanup interval",
	)
	flag.StringVar(
		&f.Redis.Host,
		"redisHost",
		defaults.Redis.Host,
		"Redis Host",
	)
	flag.UintVar(
		&redisPort,
		"redisPort",
		uint(defaults.Redis.Port),
		"Redis Port",
	)
	flag.StringVar(
		&f.Redis.Password,
		"redisPassword",
		defaults.Redis.Password,
		"Redis Password",
	)
	flag.Parse()

	if f.IP == "" {
		f.IP = defaults.IP
	}

	if serverPort > math.MaxUint16 {
		fmt.Println("Error: Server port value exceeds the range of uint16")
	}
	f.Port = uint16(serverPort)

	if f.Env == "" {
		f.Env = defaults.Env
	}

	if f.Redis.Host == "" {
		f.Redis.Host = defaults.Redis.Host
	}

	if redisPort > math.MaxUint16 {
		fmt.Println("Error: Redis port value exceeds the range of uint16")
	}
	f.Redis.Port = uint16(redisPort)

	if f.Redis.Password == "" {
		f.Redis.Password = defaults.Redis.Password
	}

	return f
}
