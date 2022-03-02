package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Request struct {
	Id   int
	Time time.Time
}

var Queue = make(chan Request, 100) //Queue

func main() {
	fmt.Println("Testing")
	http.HandleFunc("/queue", Router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Router(w http.ResponseWriter, r *http.Request) {
	go func() { // Enqueue

		for j := 1; j < 10; j++ {
			Reqs := Request{Id: j, Time: time.Now()}
			Queue <- Reqs
		}
		close(Queue)
		fmt.Println("Written into the queue")

	}()

	go func() { //Dequeue
		for i := range Queue {
			fmt.Println("Request is -->", i)
		}
		fmt.Println("Read from queue")

	}()
	return
}
