package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"db"`
	CloudType string `json:"cloud_Type"`
}

func InitConfig() Config {
	var config Config
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}

func InitDB(config Config) *gorm.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s application_name=compaction_fix sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Password)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
