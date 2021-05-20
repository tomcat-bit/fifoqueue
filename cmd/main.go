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
		q.Insert(fmt.Sprintf("kake %d", i))
		log.Println(q.Back())
	}

}