package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"example.com/m/config"
)

func main() {
	cfg := config.Load(".")
	// fmt.Println(cfg)
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Mysql.User,     // Foydalanuvchi nomi
		cfg.Mysql.Password, // Parol
		cfg.Mysql.Host,     // Host (masalan, "localhost")
		cfg.Mysql.Port,     // Port (masalan, "3306")
		cfg.Mysql.Database, // Ma'lumotlar bazasi nomi
	)

	mysqlConn, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	defer mysqlConn.Close() // Dastur tugagach ulanishni yopish

	// Ulanishni tekshirish
	err = mysqlConn.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	} else {
		log.Println("Connection sucss")
	}

}
