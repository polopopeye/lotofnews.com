package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username string = "darkfurius"
const password string = "7e7abetadox_K@"
const host string = "localhost"
const port int = 3306
const database string = "golotofnews"

func CreateConnection() {
	if connection, err := sql.Open("mysql", generateURL()); err != nil {
		panic(err)
	} else {
		db = connection
		fmt.Println("Conexión a DB exitosa")
	}
}
func ExistTable(tableName string) bool {
	//SHOW TABLES LIKE 'newsorigin'
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	return rows.Next()
}
func CreateTable(tableName, schema string) { //dos valores string Nombre de la tabla a crear y datos.
	if !ExistTable(tableName) {
		_, err := db.Exec(schema)
		if err != nil {
			log.Println(err)
		}
	}
}
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
func CloseConnection() {
	db.Close()
}

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, database)
} //db, err := sql.Open("mysql", "user:password@/dbname")
//<username:<password>@tcp
