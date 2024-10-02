package config

import (
	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	LimitIp    int `mapstructure:"LIMITE_IP"`
	LimitToken int `mapstructure:"LIMITE_TOKEN"`
	BlockTime  int `mapstructure:"BLOCK_TIME_IN_SECONDS"`
}

func GetConfig() *Config {
	return cfg
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("app_config") //nome das configurações
	viper.SetConfigType("env")        //formato do arquivo de configuração
	viper.AddConfigPath(path)         //caminho do arquivo de configuração
	viper.SetConfigFile(".env")       //nome do arquivo de config

	//Lê as variáveis de ambiente
	err := viper.ReadInConfig()
	if err != nil {
		//return nil, err
		panic(err) // se não conseguir carregar as variáveis de ambiente, vamos lançar um panic
	}

	//Passa os dados lidos para a struct
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
