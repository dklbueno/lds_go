package utils

import (
	"os"
)

type Config struct {
	Mysql         bool
	MysqlDbUser   string
	MysqlDbPass   string
	MysqlDbName   string
	MysqlDbHost   string
	MysqlDbPort   string
	MysqlDbTable  string
	MysqlDbColumn string

	Postgres         bool
	PostgresDbUser   string
	PostgresDbPass   string
	PostgresDbName   string
	PostgresDbHost   string
	PostgresDbPort   string
	PostgresDbTable  string
	PostgresDbColumn string

	Newrelic           bool
	NewrelicEntity     string
	NewrelicAppName    string
	NewrelicLicenseKey string
	ChannelCap         int
	QtdWorker          int
}

var env *Config

func (conf *Config) setEnv() {
	// env.DbConnection := "mysql"
	// env.DbHost := "10.220.1.7"
	// env.DbPort := "3306"
	// env.DbName := "movida"
	// env.DbUser := "movidarac"
	// env.DbPass := "jG6CCp2@Fa3Vek8t[)YPdXgG"

	// env.DbConnection := "postgres"
	// env.DbHost := "127.0.0.1"
	// env.DbPort := "5432"
	// env.DbName := "lds"
	// env.DbUser := "postgres"
	// env.DbPass := "123456"

	//LicenseKey : 2f09561f5089d3ed2218d2dfcb0bf0335c692e7e
	//AppName : OTA_Metrics
	conf.Mysql = StringToBool(os.Getenv("mysql"))
	conf.MysqlDbHost = os.Getenv("mysql_db_host")
	conf.MysqlDbPort = os.Getenv("mysql_db_port")
	conf.MysqlDbName = os.Getenv("mysql_db_name")
	conf.MysqlDbUser = os.Getenv("mysql_db_user")
	conf.MysqlDbPass = os.Getenv("mysql_db_pass")
	conf.MysqlDbTable = os.Getenv("mysql_db_table")
	conf.MysqlDbColumn = os.Getenv("mysql_db_column")

	conf.Postgres = StringToBool(os.Getenv("postgres"))
	conf.PostgresDbHost = os.Getenv("postgres_db_host")
	conf.PostgresDbPort = os.Getenv("postgres_db_port")
	conf.PostgresDbName = os.Getenv("postgres_db_name")
	conf.PostgresDbUser = os.Getenv("postgres_db_user")
	conf.PostgresDbPass = os.Getenv("postgres_db_pass")
	conf.PostgresDbTable = os.Getenv("postgres_db_table")
	conf.PostgresDbColumn = os.Getenv("postgres_db_column")

	conf.Newrelic = StringToBool(os.Getenv("newrelic"))
	conf.NewrelicEntity = os.Getenv("newrelic_entity")
	conf.NewrelicAppName = os.Getenv("newrelic_app_name")
	conf.NewrelicLicenseKey = os.Getenv("newrelic_license_key")
	conf.ChannelCap = StringToInt(os.Getenv("channel_cap"))
	conf.QtdWorker = StringToInt(os.Getenv("qtd_worker"))

}

func GetEnv() *Config {
	if env == nil {
		env = new(Config)
		env.setEnv()
	}
	return env
}
