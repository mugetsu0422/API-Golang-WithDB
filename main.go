package main

import(
	"fmt"
	// "net/http"
	// "github.com/labstack/echo/v4"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	err := orm.RegisterDataBase("default", "mysql", "root:123456@/echotest?charset=utf8")
	if err != nil{
		fmt.Println("Failed to connect database", err)
	}

	// Database alias
	name := "default"

	// Drop table and re-create
	force := true

	// Print log
	verbose := true

	// Error
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil{
		fmt.Println("Failed to run sync", err)
	}
}

func main() {
	fmt.Println("Hello, world.")
}