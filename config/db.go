package config

import(
	"fmt"
)

const (
	DBUser = "root"
	DBPassword = "123456"
	DBName = "musicapp"
	DBHost = "localhost"
	DBPort = "3306"
	DBType = "mysql"
)

func GetDBType() string {
	return DBType
}

func GetMySQLConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
	DBUser, DBPassword, DBHost, DBPort, DBName)
	return dataBase
}