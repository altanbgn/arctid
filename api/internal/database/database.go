package database

import (
  "log"

  "arctid/api/internal/config"
  "arctid/api/internal/models"

  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var DB *gorm.DB

func Connect() {
  var err error

  DB, err = gorm.Open(postgres.Open(config.Env.DB_URL), &gorm.Config{
    SkipDefaultTransaction: true,
    PrepareStmt: true,
  })
  if err != nil {
    log.Fatalln("Error connecting to database")
  }

  err = DB.AutoMigrate(
    &models.User{},
  )
  if err != nil {
    log.Fatalln("Error migrating database")
  }

  log.Println("Connected to database")
}
