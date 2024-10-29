package registry

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func (r *registry) GetConnection() DbModel {
	model := DbModel{}

	dbUser, err := r.conf.String("database.dbUser")
	if err != nil {
		panic(err)
	}
	dbPass, err := r.conf.String("database.dbPass")
	if err != nil {
		panic(err)
	}
	dbHost, err := r.conf.String("database.dbHost")
	if err != nil {
		panic(err)
	}
	dbPort, err := r.conf.String("database.dbPort")
	if err != nil {
		panic(err)
	}
	dbName, err := r.conf.String("database.dbName")
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
func (r *registry) NewUrlMappingDatabaseClient() (db *gorm.DB, err error) {
	if ormClient != nil {
		return ormClient, nil
	}

	model := r.GetConnection()
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

	// ormClient.Logger.LogMode(logger.Silent)

	return ormClient, nil
}
