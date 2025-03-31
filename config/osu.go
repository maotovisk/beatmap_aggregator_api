package config

type OsuConfig struct {
	ClientID     string `env:"OSU_CLIENT_ID"`
	ClientSecret string `env:"OSU_CLIENT_SECRET"`
}
