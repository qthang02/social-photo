package common

import "github.com/spf13/viper"

type Config struct {
	DB_CONN_STR    string `mapstructure:"DB_CONN_STR"`
	S3APIKey       string `mapstructure:"S3APIKEY"`
	S3APISecret    string `mapstructure:"S3APISECRET"`
	S3Region       string `mapstructure:"S3REGION"`
	S3BucketName   string `mapstructure:"S3BUCKETNAME"`
	S3Domain       string `mapstructure:"S3DOMAIN"`
	JWT_SECRET_KEY string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
