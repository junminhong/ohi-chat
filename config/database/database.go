package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/junminhong/ohi-chat/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBINFO struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbName     string
	dbPort     string
}

func (dbInfo DBINFO) setupDBInfo() *DBINFO {
	err := godotenv.Load()
	if err != nil {
		utils.ShowErrSystemMsg("Failed to load env file")
	}
	dbInfo.dbUser = os.Getenv("DB_USER")
	fmt.Println(os.Getenv("DB_PASSWORD"))
	dbInfo.dbPassword = os.Getenv("DB_PASSWORD")
	dbInfo.dbHost = os.Getenv("DB_HOST")
	dbInfo.dbName = os.Getenv("DB_NAME")
	dbInfo.dbPort = os.Getenv("DB_PORT")
	return &dbInfo
}

func GetDB() *gorm.DB {
	dbInfo := &DBINFO{}
	dbInfo = dbInfo.setupDBInfo()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		dbInfo.dbHost, dbInfo.dbUser, dbInfo.dbPassword, dbInfo.dbName, dbInfo.dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		utils.ShowErrSystemMsg("Failed to connect DB")
	}
	return db
}
