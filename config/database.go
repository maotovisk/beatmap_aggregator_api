package config

type SqliteConfig struct {
	Path string `env:"DATABASE_PATH"`
}
