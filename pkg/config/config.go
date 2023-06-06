package config

import (
	"os"
)

var (
	Conf struct {
		Mysql struct {
			Host     string
			Port     string
			User     string
			Password string
			DbName   string
		}
		Env string
	}
)

func init() {
	Conf.Env = os.Getenv("Env")

	switch Conf.Env {
	case "prod":
		{
			//internal
			Conf.Mysql.Host = os.Getenv("DbHost")
			Conf.Mysql.Port = os.Getenv("DbPort")
			Conf.Mysql.User = os.Getenv("DbId")
			Conf.Mysql.Password = os.Getenv("DbPwd")
			Conf.Mysql.DbName = os.Getenv("DbName")

			//external

		}
		break

	case "test":
		{
			//internal
			Conf.Mysql.Host = os.Getenv("DbHost")
			Conf.Mysql.Port = os.Getenv("DbPort")
			Conf.Mysql.User = os.Getenv("DbId")
			Conf.Mysql.Password = os.Getenv("DbPwd")
			Conf.Mysql.DbName = os.Getenv("DbName")

			//external

		}
		break

	default:
		{
			//external
			//Conf.Mysql.Host = "121.40.64.95"
			//Conf.Mysql.Port = "3306"
			//Conf.Mysql.User = ""
			//Conf.Mysql.Password = ""
			//Conf.Mysql.DbName = "userdb"

			//LCD-local
			Conf.Mysql.Host = "121.40.64.95"
			Conf.Mysql.Port = "3306"
			Conf.Mysql.User = "liupeiyu"
			Conf.Mysql.Password = "123456"
			Conf.Mysql.DbName = "dynamic"
		}
	}
}
