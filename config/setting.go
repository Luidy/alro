package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// DefaultConfig ...
type DefaultConfig struct {
	ConfENV        string
	ConfServerPort int
	ConfDBHOST     string
	ConfDBPORT     int
	ConfDBNAME     string
	ConfDBUSER     string
	ConfDBPASS     string
}

var defaultConfig = DefaultConfig{
	ConfServerPort: 10500,
	ConfDBHOST:     "alro",
	ConfDBPORT:     3306,
	ConfDBNAME:     "alrodb",
	ConfDBUSER:     "alrouser",
	ConfDBPASS:     "alropass",
}

// Alro ...
var Alro *viper.Viper

func init() {
	pflag.IntP("port", "p", defaultConfig.ConfServerPort, "server port")
	pflag.String("db_host", defaultConfig.ConfDBHOST, "db host")
	pflag.Int("db_port", defaultConfig.ConfDBPORT, "db port")
	pflag.String("db_name", defaultConfig.ConfDBNAME, "db name")
	pflag.String("db_user", defaultConfig.ConfDBUSER, "db user")
	pflag.String("db_pass", defaultConfig.ConfDBPASS, "db pass")

	pflag.Parse()
	var err error
	Alro, err = setConfig(map[string]interface{}{})
	if err != nil {
		fmt.Printf("Config setting Error: %v\n", err)
		os.Exit(1)
	}
	Alro.BindPFlags(pflag.CommandLine)
}

func setConfig(defaults map[string]interface{}) (*viper.Viper, error) {
	// setting config process: default -> env file -> command line
	v := viper.New()
	for k, d := range defaults {
		v.SetDefault(k, d)
	}

	v.AddConfigPath("./")
	v.AddConfigPath("./config")
	v.SetConfigName(".env")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}
