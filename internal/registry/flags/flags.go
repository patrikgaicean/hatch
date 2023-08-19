package flags

import (
	"flag"
	"fmt"
	"math"
)

type flags struct {
	Port  uint16
	Env   string
	Redis struct {
		Host     string
		Port     uint16
		Password string
	}
}

func ParseFlags() flags {
	var (
		f          flags
		serverPort uint
		redisPort  uint
	)

	flag.UintVar(&serverPort, "port", 8080, "Registry port")
	flag.StringVar(
		&f.Env,
		"env",
		"development",
		"Environment (development|staging|production)",
	)
	flag.StringVar(&f.Redis.Host, "redisHost", "redis-db", "Redis Host")
	flag.UintVar(&redisPort, "redisPort", 6379, "Redis Port")
	flag.StringVar(
		&f.Redis.Password,
		"redisPassword",
		"eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		"Redis Password",
	)
	flag.Parse()

	if serverPort > math.MaxUint16 {
		fmt.Println("Error: Server port value exceeds the range of uint16")
		// todo: return error
	}
	f.Port = uint16(serverPort)

	if redisPort > math.MaxUint16 {
		fmt.Println("Error: Redis port value exceeds the range of uint16")
		// todo: return error
	}
	f.Redis.Port = uint16(redisPort)

	return f
}
