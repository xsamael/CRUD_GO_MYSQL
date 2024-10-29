package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func Conn() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("db_user"),
		os.Getenv("db_password"),
		os.Getenv("db_localhost"),
		os.Getenv("db_port"),
		os.Getenv("db_name"))

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Panic(err)
	}
	Db = db
	fmt.Println("Exitosa")
}
