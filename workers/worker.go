package workers

import (
	"lds/controllers"
	. "lds/queues"
	"time"
)

func WorkerLog(queue *Queue) {

	for {
		select {
		case r := <-queue.Ch:
			//fmt.Println("Worker CreateLogOTA", r)
			controllers.CreateLogOTA(r.Payload)
		case <-time.After(50 * time.Millisecond):
			//fmt.Printf(".")
		}
	}

}
