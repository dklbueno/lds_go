package models

import (
	"database/sql"
	"fmt"
	. "lds/syslogLDS"
	. "lds/utils"
)

type DBPostgres struct {
	instancePostgres *sql.DB
	health           Health `json:"health"`
	dbError          DbError
}

func NewDBPostgres() (*DBPostgres, error) {
	env := GetEnv()
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		env.PostgresDbHost,
		env.PostgresDbUser,
		env.PostgresDbName,
		env.PostgresDbPass,
		env.PostgresDbPort)
	db := new(DBPostgres)
	var err error
	db.instancePostgres, err = sql.Open(`postgres`, dbUri)
	if err != nil {
		panic(err)
	}

	err = db.instancePostgres.Ping()
	if err != nil {
		panic(err)
	}

	return db, err
}

func (postgres *DBPostgres) Store(Info string) bool {
	sqlStatement := `
	INSERT INTO metrics(info) VALUES($1)
	RETURNING id`
	id := 0
	err := postgres.instancePostgres.QueryRow(sqlStatement, Info).Scan(&id)
	if err != nil {
		GetInstanceLog().Err(err.Error())
		//panic(err)
		return false
	}
	fmt.Println("New record ID is:", id)
	return true
}

func (postgres *DBPostgres) ShutDown() {

}

func (postgres *DBPostgres) HelthCheck() Health {
	drive := `postgres`
	status := "success"
	var err error
	if postgres.dbError.Status == true {
		err = postgres.instancePostgres.Ping()
		if err != nil {
			status = "error"
		}
	}
	var health = Health{Drive: drive, Status: status}
	return health
}

func (postgres *DBPostgres) Error() DbError {
	return postgres.dbError
}
