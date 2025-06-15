package utils

type AppConfig struct {
	ServerHost      string
	ServerPort      string
	RedisServerHost string
	RedisServerPort string
	DBHost          string
	DBPort          int
	DBUser          string
	DBPassword      string
	DBName          string
	DBSSLMode       string
}

var Config *AppConfig

func UpdateVariables() {
	Config = &AppConfig{
		ServerHost:      "localhost",
		ServerPort:      "8080",
		RedisServerHost: "default:AZtqAAIjcDFhOGIxMTlhYzI0YzM0M2UzYjE2ODg1ZjM2NGU0ZDkyYnAxMA@allowed-ray-39786.upstash.io",
		RedisServerPort: "6379",
		DBHost:          "ep-ancient-bonus-a4wkt899-pooler.us-east-1.aws.neon.tech",
		DBPort:          5432,
		DBUser:          "neondb_owner",
		DBPassword:      "npg_IzdP9VJ3QMXi",
		DBName:          "neondb",
		DBSSLMode:       "require",
	}
}
