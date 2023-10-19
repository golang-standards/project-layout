package config

import "log"

type Config struct {
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func LoadAndStorageConfig() Config {
	v := viper.New()
	v.SetEnvPrefix("SERV")
	v.SetDefault("PORT", "8080")
	v.SetDefault("DBUSER", "postgres")
	v.SetDefault("DBPASS", "")
	v.SetDefault("DBHOST", "")
	v.SetDefault("DBPORT", "5432")
	v.SetDefault("DBNAME", "postgres")

	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		log.Panic(err)
	}
	return config
}

func (config *Config) GetDBString() string {
	return fmt.SPrintf(`postgres://%s:%s@%s:%s/%s`, config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)
}
