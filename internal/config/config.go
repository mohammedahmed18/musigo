package config

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	configuration *Configuration
	configFileExt = ".json"
	configType    = "json"
)

type Configuration struct {
	Server ServerConfiguration
}

type ServerConfiguration struct {
	Env                     string
	Port                    string
	Passphrase              string
	Secret                  string
	GeneratedPasswordLength int
	SessionLifetime         int
}

// Init initializes the configuration manager
func Init(configPath, configName string) (*Configuration, error) {
	configFilePath := filepath.Join(configPath, configName) + configFileExt
	initializeConfig(configPath, configName)

	// Bind environment variables
	bindEnvs()

	// Set default values
	setDefaults()

	// Read or create configuration file
	if err := readConfiguration(configFilePath); err != nil {
		return nil, err
	}

	// Auto read env variables
	viper.AutomaticEnv()

	// Unmarshal config file to struct
	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

// read configuration from file
func readConfiguration(configFilePath string) error {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		// if file does not exist, simply create one
		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			os.Create(configFilePath)
		} else {
			return err
		}
		// let's write defaults
		if err := viper.WriteConfig(); err != nil {
			return err
		}
	}
	return nil
}

// initialize the configuration manager
func initializeConfig(configPath, configName string) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
}

func bindEnvs() {
	viper.BindEnv("server.env", "MG_ENV")
	viper.BindEnv("server.port", "PORT")
	viper.BindEnv("server.passphrase", "MG_SERVER_PASSPHRASE")
	viper.BindEnv("server.secret", "MG_SERVER_SECRET")
	viper.BindEnv("server.sessionLifetime", "MG_SESSION_LIFETIME")
	viper.BindEnv("server.generatedPasswordLength", "MG_SERVER_GENERATED_PASSWORD_LENGTH")
}

func setDefaults() {
	viper.SetDefault("server.env", "prod")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.passphrase", generateKey())
	viper.SetDefault("server.secret", generateKey())
	viper.SetDefault("server.generatedPasswordLength", 16)
	viper.SetDefault("server.sessionLifetime", "1d")

}

func generateKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "add-your-key-to-here"
	}
	keyEnc := base64.StdEncoding.EncodeToString(key)
	return keyEnc
}
