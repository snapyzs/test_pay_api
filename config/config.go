package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port       string
	Host       string
	PortDB     string
	HostDB     string
	NameDB     string
	UsernameDB string
	PasswordDB string
	SSLModeDB  string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	cfg := &Config{
		Port:       os.Getenv("PORT"),
		Host:       os.Getenv("HOST"),
		PortDB:     os.Getenv("PORT_DB"),
		HostDB:     os.Getenv("HOST_DB"),
		NameDB:     os.Getenv("NAME_DB"),
		UsernameDB: os.Getenv("USERNAME_DB"),
		PasswordDB: os.Getenv("PASSWORD_DB"),
		SSLModeDB:  os.Getenv("SSL_MODE"),
	}
	return cfg
}
