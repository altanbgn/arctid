package database

import (
  "log"

  "arctid/api/configs"

  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var Database *gorm.DB

func Connect() {
  configs := configs.Get()

  db, err := gorm.Open(postgres.Open(configs.DB_URL), &gorm.Config{
    SkipDefaultTransaction: true,
    PrepareStmt: true,
  })

  if err != nil {
    log.Fatalln("Error connecting to database")
  }

  Database, err := db.DB()

  if err != nil {
    log.Fatalln("Error connecting to database")
  }

  Database.SetMaxIdleConns(configs.DB_MAX_IDLE_CONNS)
  Database.SetMaxOpenConns(configs.DB_MAX_OPEN_CONNS)
  Database.SetConnMaxLifetime(configs.DB_CONN_MAX_LIFETIME)

  log.Println("Connected to database")
}
