package configs

import (
  "os"
  "log"
  "strconv"
  "time"

  "github.com/joho/godotenv"
)

type Config struct {
  HOST string
  PORT int
  TIMEOUT time.Duration
  DEBUG bool

  JWT_SECRET_KEY string
  JWT_EXPIRE int

  DB_URL string
  DB_MAX_OPEN_CONNS int
  DB_MAX_IDLE_CONNS int
  DB_CONN_MAX_LIFETIME time.Duration
}

var config = &Config{}

func Get() *Config {
  return config;
}

func LoadDotenv () {
  err := godotenv.Load()
  if (err != nil) {
    log.Fatal("Error loading .env file")
  }

  log.Println("Loaded .env file")

  config.HOST = os.Getenv("HOST")
  config.PORT, _ = strconv.Atoi(os.Getenv("PORT"))
  config.DEBUG, _ = strconv.ParseBool(os.Getenv("DEBUG"))
  timeout, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
  config.TIMEOUT = time.Duration(timeout) * time.Second

  config.JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
  config.JWT_EXPIRE, _ = strconv.Atoi(os.Getenv("JWT_EXPIRE"))

  config.DB_URL = os.Getenv("DB_URL")
  config.DB_MAX_OPEN_CONNS, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
  config.DB_MAX_IDLE_CONNS, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
  lifetime, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))
  config.DB_CONN_MAX_LIFETIME = time.Duration(lifetime) * time.Second
}
