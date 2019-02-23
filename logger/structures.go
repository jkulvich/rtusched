package logger

type Config struct {
	Level string `yaml:"level"`
	Format string `yaml:"format"`
	CallerInfo bool `yaml:"callerInfo"`
}