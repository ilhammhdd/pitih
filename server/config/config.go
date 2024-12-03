package config

import "time"

const SystemSuperAdminPasswordEnvVar = "SYSTEM_SUPERADMIN_PASSWORD"

type config struct {
	System struct {
		Name string `mapstructure:"name"`
		Http struct {
			Timeout struct {
				Read             uint16 `mapstructure:"read"`
				GracefulShutdown uint16 `mapstructure:"graceful-shutdown"`
			} `mapstructure:"timeout"`
			Host string `mapstructure:"host"`
			Port string `mapstructure:"port"`
		} `mapstructure:"http"`
		TokenStoreFilename string        `mapstructure:"token-store-file-name"`
		TokenTTLSecond     time.Duration `mapstructure:"token-ttl-second"`
		TokenKey           string        `mapstructure:"token-key"`
		SuperAdminPassword string        `mapstructure:"superadmin-password"`
	} `mapstructure:"system"`
	RelationalDB struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"relational-database"`
}

var Config *config

func ToSecondDuration[T int64 | int | uint16](t T) time.Duration {
	return time.Duration(t) * time.Second
}
