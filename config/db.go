package config

import(
	"fmt"
)

const (
	DBUser = "user"
	DBPassword = "user"
	DBName = "mydb"
	DBHost = "localhost"
	DBPort = "3306"
	DBType = "mysql"
)

func GetDBType() string {
	return DBType
}

func GetMySQLConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
	DBUser, DBPassword, DBHost, DBPort, DBType)
	return dataBase
}