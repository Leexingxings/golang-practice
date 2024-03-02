package config

type Config struct {
	JWTSecret string `json:"jwt_secret" env:"JWT_SECRET"`
}
