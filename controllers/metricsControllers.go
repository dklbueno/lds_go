package controllers

import (
	. "lds/models"
)

// CreateLogOTA cria um novo log
func CreateLogOTA(stringJson string) {
	//dbs, _ := NewDBs()
	StoreAll(stringJson)
}
