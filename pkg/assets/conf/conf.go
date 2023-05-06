package conf

import "github.com/spf13/viper"

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (Config, error) {
	viper.AddConfigPath("./pkg/assets/conf/envs")
	viper.SetConfigName("dev")
	viper.SetConfigFile("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		return Config{}, err
	}
	return conf, nil
}
