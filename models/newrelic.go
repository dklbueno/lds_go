package models

import (
	"encoding/json"
	. "lds/utils"
	"time"

	newrelic "github.com/newrelic/go-agent"
)

type DBNewRelic struct {
	instanceNewRelic newrelic.Application
	health           *Health `json:"health"`
	dbError          DbError
}

func NewDBNewRelic() (*DBNewRelic, error) {
	env := GetEnv()

	db := new(DBNewRelic)

	var err error

	config := newrelic.NewConfig(env.NewrelicAppName, env.NewrelicLicenseKey)
	db.instanceNewRelic, err = newrelic.NewApplication(config)
	if nil != err {
		panic(err)
	}

	// Wait for the application to connect.
	if err = db.instanceNewRelic.WaitForConnection(10 * time.Second); nil != err {
		panic(err)
	}

	return db, err
}

func (nrelic *DBNewRelic) Store(Info string) bool {
	env := GetEnv()
	var payload map[string]interface{}
	json.Unmarshal([]byte(Info), &payload)

	nrelic.instanceNewRelic.RecordCustomEvent(env.NewrelicEntity, payload)

	return true
}

func (nrelic *DBNewRelic) ShutDown() {

}

func (nrelic *DBNewRelic) HelthCheck() Health {
	drive := `newrelic`
	status := "success"
	if nrelic.dbError.Status == true {
		status = "error"
	}
	var health = Health{Drive: drive, Status: status}
	return health
}

func (nrelic *DBNewRelic) Error() DbError {
	return nrelic.dbError
}
