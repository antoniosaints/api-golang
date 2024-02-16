package db

import (
	"api/configs"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetConfig()
	sc := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.Host, conf.Port, conf.Username, conf.DbName, conf.Password)
	conn, err := sql.Open("mysql", sc)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()

	return conn, err
}
