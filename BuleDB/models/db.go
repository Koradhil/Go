package models

import (
    "database/sql"
    _ "github.com/lib/pq"
//	  _ "github.com/go-sql-driver/mysql"
    "log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Println("Something is wrong with the database")
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Println("Ping failed")
        log.Panic(err)
    } else {
      log.Println("DB connection running")
    }
}
