package httpserver

type Config struct {
	Port                      int    `mapstructure:"PORT"`
	LogLevel                  string `mapstructure:"LOG_LEVEL"`
	ReadTimeoutMs             int    `mapstructure:"READ_TIMEOUT_MS"`
	WriteTimeoutMs            int    `mapstructure:"WRITE_TIMEOUT_MS"`
	GracefulShutdownTimeoutMs int    `mapstructure:"GRACEFUL_SHUTDOWN_TIMEOUT_MS"`
}
