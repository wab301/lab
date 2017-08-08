package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main1() {
	db, err := sql.Open("mysql", "game:ftrunk@tcp(127.0.0.1:3306)/game_admin?charset=utf8mb4")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("SELECT `db_host_name`,`db_port`,`db_user`,`db_pwd`,`db_name` FROM `server` WHERE `server_id` = ?")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	var (
		dbHostName string
		dbPort     uint16
		dbUser     string
		dbPwd      string
		dbName     string
	)
	err = stmt.QueryRow(2).Scan(&dbHostName, &dbPort, &dbUser, &dbPwd, &dbName)
	println(dbHostName, dbPort, dbUser, dbPwd, dbName)
}
func main() {

	adminDB, err := New("game:ftrunk@tcp(127.0.0.1:3306)/game_admin?charset=utf8mb4", map[string]string{
		"select":           "select `db_name` from `server` where `server_id`=?",
		"selectServerInfo": "SELECT `db_host_name`,`db_port`,`db_user`,`db_pwd`,`db_name` FROM `server` WHERE `server_id` = ?",
	})
	if err != nil {
		log.Fatalf("admin params error: %v", err)
	}

	// var (
	// 	dbHostName string
	// 	dbPort     int16
	// 	dbUser     string
	// 	dbPwd      string
	// 	dbName     string
	// )

	// err = adminDB.QueryRow("selectServerInfo", 2).Scan(&dbHostName, &dbPort, &dbUser, &dbPwd, &dbName)
	// if err != nil {
	// 	log.Fatal("==========", err.Error())

	// }
	// log.Fatal("====", dbHostName, dbPort, dbUser, dbPwd, dbName)
	var name string
	err = adminDB.QueryRow("select", 1).Scan(&name)
	panic(err)
}
