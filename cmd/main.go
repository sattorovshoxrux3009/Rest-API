package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"example.com/m/config"
	"example.com/m/storage"
	"example.com/m/storage/repo"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/google/uuid"
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

	/*id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Error generating UUID: ", err)
	}
	strg := storage.NewStorage(mysqlConn)
	user, err := strg.User().Create(context.TODO(), &repo.User{
		ID:        id.String(),
		FirstName: "Satorov",
		LastName:  "Shohruh",
		Email:     "Sattorovshohruh300s9@gmaisl.coma",
		Password:  "12345678a",
	})
	if err != nil {
		log.Fatal("Error creating user: ", err)
	}
	fmt.Println(user)

	strg := storage.NewStorage(mysqlConn)
	userGet, err := strg.User().Get(context.TODO(), "1d5021e5-ed05-4cc8-af1c-5c574515a17c")
	if err != nil {
		log.Fatal("Error getting user: ", err)
	}
	fmt.Println(userGet)*/

	strg := storage.NewStorage(mysqlConn)
	err = strg.User().Update(context.TODO(), &repo.UpdateUser{
		ID:        "b35bf967-b9e3-4982-bad6-6f23c39ec9c0",
		FirstName: "XSSSShohruh",
		LastName:  "xsssSatorov",
	})
	if err != nil {
		log.Fatal("Error updating user: ", err)
	}	
}
