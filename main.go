package main

import (
	"fmt"
	"lds/controllers"
	. "lds/queues"
	. "lds/utils"
	"net/http"

	. "lds/models"

	. "lds/syslogLDS"

	. "lds/workers"

	"github.com/gorilla/mux"
)

func InitApp() {

	env := GetEnv()
	fmt.Println("env", env)
	queue := GetInstanceQueue()
	fmt.Println("Queue main", queue)
	dbs, errDbs := NewDBs()
	fmt.Println("DBs", dbs)

	if len(errDbs) > 0 {
		for _, db := range dbs {
			GetInstanceLog().Notice(db.Error().Message)
			fmt.Println(db.Error().Message)
		}
	}

	for i := 0; i <= env.QtdWorker; i++ {
		go WorkerLog(queue)
	}
}

func LogHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("LogHandler")

	if ValidateJson(w, r) {
		queue := GetInstanceQueue()

		var request = RequestQueue{StringJson}
		queue.Enqueue(request)
	}

	//w.Write([]byte(StringJson))

}

// função principal para executar a api
func main() {

	InitApp()

	router := mux.NewRouter()
	router.HandleFunc("/log", LogHandler).Methods("POST")
	router.HandleFunc("/healthcheck", controllers.Healthcheck).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8080", router))
	err := http.ListenAndServe(":8080", router)
	GetInstanceLog().Err(err.Error())
}
