package models

import (
	"database/sql"
	"fmt"
	. "lds/syslogLDS"
	. "lds/utils"
)

type DBMysql struct {
	instanceMysql *sql.DB
	health        *Health `json:"health"`
	dbError       DbError
}

func NewDBMysql() (*DBMysql, error) {
	env := GetEnv()
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		env.MysqlDbUser,
		env.MysqlDbPass,
		env.MysqlDbHost,
		env.MysqlDbPort,
		env.MysqlDbName)
	db := new(DBMysql)
	var err error
	db.instanceMysql, err = sql.Open(`mysql`, dbUri)

	if err != nil {
		//panic(err)
		db.dbError.Status = true
		db.dbError.Message = err.Error()
		return db, err
	}

	err = db.instanceMysql.Ping()
	if err != nil {
		//panic(err)
		db.dbError.Status = true
		db.dbError.Message = err.Error()
		return db, err
	}

	return db, err
}

func (mysql *DBMysql) Store(Info string) bool {
	sqlStatement := "INSERT INTO metrics(info) VALUES(?)"
	query, err := mysql.instanceMysql.Prepare(sqlStatement)
	if err != nil {
		GetInstanceLog().Err(err.Error())
		// panic(err.Error)
		return false
	}

	res, err := query.Exec(Info)
	if err != nil {
		GetInstanceLog().Err(err.Error())
		// panic(err.Error)
		return false
	}

	id, _ := res.LastInsertId()

	fmt.Println("New record ID is:", id)

	return true
}

func (mysql *DBMysql) ShutDown() {

}

func (mysql *DBMysql) HelthCheck() Health {
	drive := `mysql`
	status := "success"
	var err error
	if mysql.dbError.Status == true {
		err = mysql.instanceMysql.Ping()
		if err != nil {
			status = "error"
		}
	}
	var health = Health{Drive: drive, Status: status}
	return health
}

func (mysql *DBMysql) Error() DbError {
	return mysql.dbError
}
