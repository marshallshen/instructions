package app

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "gopkg.in/gorp.v1"
  "log"
)

func initDb() *gorp.DbMap {
  db, err := sql.Open("mysql", "root:password@/instructions")
  checkErr(err, "sql.Open failed")

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
  dbmap.AddTableWithName(Instruction{}, "Instruction").SetKeys(true, "Id")

  err = dbmap.CreateTablesIfNotExists()
  checkErr(err, "Create table failed")
  return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
    log.Fatalln(msg, err)
  }
}
