
package fifoqueue 

import (
	"testing"
	"fmt"
	"container/list"
	//"github.com/tomcat-bit/fifoqueue"
)

func TestNewFIFOQueue(t *testing.T) {
	q, err := NewFIFOQueue(3)	
	if err != nil {
		t.Errorf("Got an error: %s\n", err.Error())
	}

	if q == nil {
		t.Errorf("New queue was nil\n")
	}
}

func TestInsert(t *testing.T) {
	n := 10
	q, err := NewFIFOQueue(n)	
	if err != nil {
		t.Errorf("Got an error: %s\n", err.Error())
	}

	// Insert n elements. Make sure we get what we insert
	for i := 0; i < n; i++ {
		elem := fmt.Sprintf("Element nr %d", i)
		q.Insert(elem)

		v := q.Back().(list.Element)
		fmt.Println(v)
	} 
}