package config

type Config struct {
	ListenOn string `yaml:"ListenOn"`
	RedisUrl string `yaml:"RedisUrl"`
}
