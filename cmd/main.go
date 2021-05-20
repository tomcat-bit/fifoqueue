package main

import (
	"log"
	"fmt"
	"github.com/tomcat-bit/fifoqueue"
)

// Simple example use of fifoqueue

func main() {
	q, err := fifoqueue.NewFIFOQueue(0)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		elem := fmt.Sprintf("kake %d", i)
		q.Insert(elem)
		log.Println(q.Exists(elem))
	}

	log.Println(q.Exists("k"))

}