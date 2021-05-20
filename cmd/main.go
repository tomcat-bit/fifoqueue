package main

import (
	"log"
	"fmt"
	"github.com/tomcat-bit/fifoqueue"
)

// Simple example use of fifoqueue

func main() {
	q, err := fifoqueue.NewFIFOQueue(3)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		elem := fmt.Sprintf("kake %d", i)
		q.Insert(elem)
		log.Println(q.Exists(elem))
	}

	for _, e := range q.Elements() {
		log.Println(e)	
	}

	log.Println(q.Exists("k"))

}