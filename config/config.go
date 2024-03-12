package config

type Config struct {
	JWTSecret   string `json:"jwt_secret" env:"JWT_SECRET"`
	RedisConfig redisConfig
}

type redisConfig struct {
	Host     string `json:"redis_host"`
	Port     int    `json:"redis_port"`
	Username string `json:"redis_username"`
	Password string `json:"redis_password"`
	Db       uint8  `json:"redis_db"`
}
