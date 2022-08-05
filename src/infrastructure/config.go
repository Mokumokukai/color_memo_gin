package infrastructure

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Port     string
			DBName   string
			UserName string
			Password string
		}
		Test struct {
			Host     string
			Port     string
			DBName   string
			UserName string
			Password string
		}
	}

	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)
	err := godotenv.Load(".env")
	if err != nil {
		panic("cannot load .env")
	}

	c.DB.Production.Host = "mysql_host"
	c.DB.Production.Port = "3306"
	c.DB.Production.DBName = os.Getenv("MYSQL_DATABASE")
	c.DB.Production.Password = os.Getenv("MYSQL_PASSWORD")
	c.DB.Production.UserName = os.Getenv("MYSQL_USER")

	c.Routing.Port = ":8080"
	return c
}
