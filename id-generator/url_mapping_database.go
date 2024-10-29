package main

import (
	"fmt"
	configs "github.com/joniaranguri/meli-urlshortener-challenge/id-generator/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DbModel struct {
	User     string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"dbName"`
}

func (db *DbModel) CreateConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}

var ormClient *gorm.DB

func GetConnection() DbModel {
	model := DbModel{}

	dbUser, err := configs.Conf.String("database.dbUser")
	if err != nil {
		panic(err)
	}
	dbPass, err := configs.Conf.String("database.dbPass")
	if err != nil {
		panic(err)
	}
	dbHost, err := configs.Conf.String("database.dbHost")
	if err != nil {
		panic(err)
	}
	dbPort, err := configs.Conf.String("database.dbPort")
	if err != nil {
		panic(err)
	}
	dbName, err := configs.Conf.String("database.dbName")
	if err != nil {
		panic(err)
	}
	model = DbModel{
		User:     dbUser,
		Password: dbPass,
		Host:     dbHost,
		Port:     dbPort,
		Name:     dbName,
	}
	return model
}
func NewUrlMappingDatabaseClient() (db *gorm.DB, err error) {
	if ormClient != nil {
		return ormClient, nil
	}

	model := GetConnection()
	dsn := model.CreateConnectionString()

	ormClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := ormClient.DB()
	if err != nil {
		return nil, err
	}

	// Configure connection pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	ormClient.Logger.LogMode(logger.Silent)

	return ormClient, nil
}
