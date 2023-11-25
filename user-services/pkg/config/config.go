package config

import (
	"os"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", // Database
}

var config Config // for create instance

func LoadCOnfig() (Config, error) {

	v := viper.New()

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return config, err
	}

	//Viper setup
	v.AddConfigPath(wd)
	v.SetConfigFile(".env")
	v.ReadInConfig() // for read config

	v.SetEnvPrefix("user-services")
	v.AutomaticEnv()

	for _, env := range envs {
		if err := v.BindEnv(env); err != nil {
			return config, err // error handling
		}
	}

	// then unmarshel the viper into config variable

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	// atlast validate the config file using validator pakage
	// create instance and direct validte
	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	//successfully loaded the env values into struct config

	return config, nil
}

func GetCofig() Config {
	return config

}
