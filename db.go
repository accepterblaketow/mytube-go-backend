
package main

import (
    "database/sql"
    "log"
    _ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("sqlite", "./videos.db")
    if err != nil {
        log.Fatal(err)
    }

    createTable := `
    CREATE TABLE IF NOT EXISTS videos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        description TEXT,
        tags TEXT,
        url TEXT
    );
    `
    _, err = db.Exec(createTable)
    if err != nil {
        log.Fatal(err)
    }

}
