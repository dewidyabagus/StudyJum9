package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	Env     string `mapstructure:"ENV"`
}

func LoadConfig(files ...string) Config {
	/*
		Load file config
		1. Menentukan lokasi file
		2. Menentukan nama file
		3. Menentukan type file config | viper: JSON, ENV, TOML, YAML
		4. Load data dari file config
			a. Pengecualian environment yang sudah tersedia dalam lingkungan system
			b. load data
		5. Binding data ke dalam struct
	*/
	for _, file := range files {
		viper.AddConfigPath(file)
	}
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Read in cofig:", err.Error())
	}

	var config = Config{}
	viper.Unmarshal(&config)

	return config
}
