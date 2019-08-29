package queues

import (
	"fmt"
	. "lds/utils"
)

type RequestQueue struct {
	Payload string
}

type Queue struct {
	Ch     chan RequestQueue
	Health *HealthChan `json:"health"`
}

type HealthChan struct {
	Size int `json:"size"`
	Cap  int `json:"cap"`
}

var instance *Queue

func GetInstanceQueue() *Queue {
	if instance == nil {
		instance = new(Queue)
		instance.Ch = make(chan RequestQueue, GetEnv().ChannelCap)
		fmt.Println("GetInstance", instance)
	}
	return instance
}

func (queue *Queue) Enqueue(r RequestQueue) {
	queue.Ch <- r
	fmt.Println("Enqueue", queue.Ch, len(queue.Ch))
}

func (queue *Queue) Dequeue() RequestQueue {
	r := <-queue.Ch
	return r
}

func (queue *Queue) HelthCheck() HealthChan {
	health := HealthChan{Size: len(queue.Ch), Cap: cap(queue.Ch)}
	return health
}

// func asyncHttpGets(json []string) []*HttpResponse {
// 	ch := make(chan *HttpResponse, len(urls)) // buffered
// 	responses := []*HttpResponse{}
// 	for _, url := range urls {
// 		go func(url string) {
// 			fmt.Printf("Fetching %s \n", url)
// 			resp, err := http.Get(url)
// 			resp.Body.Close()
// 			ch <- &HttpResponse{url, resp, err}
// 		}(url)
// 	}

// 	for {
// 		select {
// 		case r := <-ch:
// 			fmt.Printf("%s was fetched\n", r.Payload)
// 			responses = append(responses, r)
// 			if len(responses) == len(urls) {
// 				return responses
// 			}
// 		case <-time.After(50 * time.Millisecond):
// 			fmt.Printf(".")
// 		}
// 	}

// 	return responses

// }
