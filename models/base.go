package models

import (
	"fmt"
	. "lds/utils"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Base interface {
	Store(string) bool
	ShutDown()
	HelthCheck() Health
	Error() DbError
}

type Health struct {
	Drive   string `json:"drive"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

var instances []Base
var instanceMysql Base
var instancePostgres Base
var instanceNewrelic Base

func NewDBs() ([]Base, []error) {
	mysql := GetEnv().Mysql
	postgres := GetEnv().Postgres
	newrelic := GetEnv().Newrelic
	var errors []error
	var errMysql error
	var errPostgres error
	var errNewrelic error
	if mysql == true && instanceMysql == nil {
		instanceMysql, errMysql = NewDBMysql()
		if errMysql != nil {
			errors = append(errors, errMysql)
			fmt.Println(errMysql.Error)
		}
		instances = append(instances, instanceMysql)
	}
	if postgres == true && instancePostgres == nil {
		instancePostgres, errPostgres = NewDBPostgres()
		if errPostgres != nil {
			errors = append(errors, errPostgres)
			fmt.Println(errPostgres.Error)
		}
		instances = append(instances, instancePostgres)
	}
	if newrelic == true && instanceNewrelic == nil {
		instanceNewrelic, errNewrelic = NewDBNewRelic()
		if errNewrelic != nil {
			errors = append(errors, errNewrelic)
			fmt.Println(errNewrelic.Error)
		}
		instances = append(instances, instanceNewrelic)
	}

	return instances, errors
}

func StoreAll(Info string) {
	if len(instances) > 0 {
		for _, instance := range instances {
			instance.Store(Info)
		}
	}
}
