package config

import (
  "os"
  "log"
  "strconv"
  "time"

  "github.com/joho/godotenv"
)

type env struct {
  HOST string
  PORT int
  TIMEOUT time.Duration
  DEBUG bool

  JWT_SECRET_KEY string
  JWT_EXPIRE int

  ARGON2_SALT string

  DB_URL string
  DB_MAX_OPEN_CONNS int
  DB_MAX_IDLE_CONNS int
  DB_CONN_MAX_LIFETIME time.Duration
}

var Env = &env{}

func LoadDotenv () {
  err := godotenv.Load()
  if (err != nil) {
    log.Fatal("Error loading .env file")
  }

  log.Println("Loaded .env file")

  Env.HOST = os.Getenv("HOST")
  Env.PORT, _ = strconv.Atoi(os.Getenv("PORT"))
  Env.DEBUG, _ = strconv.ParseBool(os.Getenv("DEBUG"))
  timeout, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
  Env.TIMEOUT = time.Duration(timeout) * time.Second

  Env.JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
  Env.JWT_EXPIRE, _ = strconv.Atoi(os.Getenv("JWT_EXPIRE"))

  Env.ARGON2_SALT = os.Getenv("ARGON2_SALT")

  Env.DB_URL = os.Getenv("DB_URL")
  Env.DB_MAX_OPEN_CONNS, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
  Env.DB_MAX_IDLE_CONNS, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
  lifetime, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))
  Env.DB_CONN_MAX_LIFETIME = time.Duration(lifetime) * time.Second
}
