package main

import (
	"fmt"
	"os"

	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/yaml"
	"simple-k8/api"
	"simple-k8/base"
	"simple-k8/log"
	"simple-k8/model"
)

type DatabaseConfig struct {
	User     string `config:"user" validate:"required"`
	Password string `config:"password" validate:"required"`
	Host     string `config:"host" validate:"required"`
	Port     int    `config:"port"`
	DbName   string `config:"dbname" validate:"required"`
}

type LogConfig struct {
	Dir        string `config:"dir" validate:"required"`
	MaxSize    int    `config:"max-logger-size"`
	MaxBackups int    `config:"max-logger-backups"`
	MaxAge     int    `config:"days-to-keep"`
}

type ApiConfig struct {
	Host     string `config:"host"`
	Port     int    `config:"port" validate:"required"`
	Restrict bool   `config:"restrict-api-check"`
}

type Config struct {
	MysqlDb DatabaseConfig `config:"mysqldb" validate:"required"`
	Log     LogConfig      `config:"log" validate:"required"`
	Api     ApiConfig      `config:"api" validate:"required"`
}

func ParseConfig(configFile string) error {
	configContent, err := yaml.NewConfigWithFile(configFile, ucfg.PathSep("."))
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found!")
		}
		return err
	}

	config := Config{}
	if err := configContent.Unpack(&config); err != nil {
		return err
	}

	clog := &config.Log
	if err := log.ConfigureLogger(clog.Dir, clog.MaxSize, clog.MaxBackups, clog.MaxAge); err != nil {
		return err
	} else {
		fmt.Printf("Saving logs at %s\n", clog.Dir)
	}

	db := &config.MysqlDb
	if err := model.ConfigureMysqlDatabase(db.Host, db.Port, db.User, db.Password, db.DbName); err != nil {
		return err
	}

	apiconf := &config.Api

	if err := base.ConfigureApiServer(apiconf.Host, apiconf.Port, &api.ApiV1Schema, apiconf.Restrict); err != nil {
		return err
	} else {
		fmt.Printf("Running API service at %d\n", apiconf.Port)
	}

	return nil
}
