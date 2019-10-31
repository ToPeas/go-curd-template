package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

type config struct {
	App struct {
		Timezone   string
		TZLocation *time.Location
		JwtSecret  string `mapstructure:"jwt_secret"`
		Port       int
		Debug      bool
		Dir        string
	}
	Database struct {
		Host      string
		Port      uint
		Username  string
		Password  string
		DBName    string `mapstructure:"dbname"`
		Charset   string `mapstructure:"charset"`
		Collation string `mapstructure:"collation"`
	}
}

var Config config

func Setup() {
	readFromFile()
}

func readFromFile() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("can not read config: %s", viper.ConfigFileUsed()))
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("viper: %w", err))
	}
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	Config.App.Dir = exPath

	tzLocation, err := time.LoadLocation(Config.App.Timezone)
	if err != nil {
		panic(err)
	}

	Config.App.TZLocation = tzLocation

	//log.Printf("%+v", Config)
}
