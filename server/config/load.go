package config

import (
	"fmt"
	"log"
	"strings"

	stdErrors "errors"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	"ilhammhdd.com/pitih/internal/pkg/errors"
)

func LoadEnv() {
	if isFileNotFound(loadEnvFile()) {
		loadEnvVar()
	}
}

func isFileNotFound(err error) bool {
	var configFileNotFoundError viper.ConfigFileNotFoundError
	return stdErrors.As(err, &configFileNotFoundError)
}

func loadEnvFile() error {
	viper.SetConfigName("env.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := readEnvFileAndSetConfig(); err != nil {
			log.Print(errors.SprintTrace(err))
		}
	})

	viper.WatchConfig()
	return readEnvFileAndSetConfig()
}

func readEnvFileAndSetConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		return errors.WrapTrace(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return errors.WrapTrace(err)
	}

	return nil
}

func loadEnvVar() {
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	envVarKeyMap := map[string]any{}
	var tempConfig config

	if err := mapstructure.Decode(tempConfig, &envVarKeyMap); err != nil {
		log.Println(errors.SprintTrace(errors.WrapTrace(err)))
		return
	}

	flatKeys, flatEnvVars := flattenKeysAndEnvVars(envVarKeyMap)
	for idx := range flatKeys {
		if err := viper.BindEnv(flatKeys[idx], flatEnvVars[idx]); err != nil {
			log.Println(errors.SprintTrace(errors.WrapTrace(err)))
			return
		}
	}

	if err := viper.Unmarshal(&tempConfig); err != nil {
		log.Println(errors.SprintTrace(errors.WrapTrace(err)))
		return
	}

	Config = &tempConfig
}

func flattenKeysAndEnvVars(keyMap map[string]any) (keys []string, envVars []string) {
	for key := range keyMap {
		var nestedKey []string
		var nestedEnvVar []string

		nestedKeyMap, ok := keyMap[key].(map[string]any)
		if ok {
			nestedKey, nestedEnvVar = flattenKeysAndEnvVars(nestedKeyMap)
		} else {
			keys = append(keys, key)
			envVars = append(envVars, strings.ToUpper(key))
		}

		for idx := range nestedKey {
			keys = append(keys, fmt.Sprintf("%s.%s", key, nestedKey[idx]))
			envVars = append(envVars, fmt.Sprintf("%s_%s", strings.ToUpper(key), nestedEnvVar[idx]))
		}
	}

	return
}
