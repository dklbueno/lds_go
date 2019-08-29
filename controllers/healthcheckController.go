package controllers

import (
	"encoding/json"
	. "lds/models"
	. "lds/queues"
	"log"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	dbs, _ := NewDBs()
	healthDBs := []Health{}
	for _, db := range dbs {
		if db.Error().Status == true {
			log.Println(db.Error().Message)
		}
		healthDBs = append(healthDBs, db.HelthCheck())
	}
	queue := GetInstanceQueue()
	//connectionDB := db.HelthCheck()
	queueHealth := queue.HelthCheck()
	health := map[string]interface{}{
		"Connection DB": healthDBs,
		"Queue":         queueHealth}

	json.NewEncoder(w).Encode(health)
}
