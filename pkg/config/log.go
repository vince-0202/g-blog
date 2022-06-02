package config

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

func DefaultLogConfig() *LogConfig {
	return &LogConfig{
		Level:    "debug",
		Filename: "./logs/zap.log",
	}
}