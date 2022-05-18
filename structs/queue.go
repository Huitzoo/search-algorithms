package structs

import (
	"sync"

	"golang.org/x/sync/semaphore"
)

//Queue is a data struct FIFO
type Queue struct {
	Size int
	data []interface{}
	mux  *sync.Mutex
	sem  semaphore.Weighted
}

func NewQueue() *Queue {
	q := &Queue{}
	q.Start()

	return q
}

//Start inizialice Queue
func (q *Queue) Start() {
	q.data = make([]interface{}, 0)
	q.Size = 0
	q.mux = &sync.Mutex{}
}

//Pop get first element of Queue
func (q *Queue) PopFIFO() interface{} {
	var data interface{}

	if q.Size == 0 {
		return nil
	}

	(*q.mux).Lock()
	if q.Size == 1 {
		data = q.data[0]
		q.data = make([]interface{}, 0)
	} else {
		data = q.data[0]
		q.data = q.data[1:]
	}

	q.Size--
	(*q.mux).Unlock()
	return data
}

func (q *Queue) PopLIFO() interface{} {
	var data interface{}

	if q.Size == 0 {
		return nil
	}

	(*q.mux).Lock()
	if q.Size == 1 {
		data = q.data[0]
		q.data = make([]interface{}, 0)
	} else {
		data = q.data[q.Size-1]
		q.data = q.data[:q.Size-1]
	}

	q.Size--
	(*q.mux).Unlock()
	return data
}

//Push put element in Queue
func (q *Queue) Push(d interface{}) {
	(*q.mux).Lock()
	q.data = append(q.data, d)
	q.Size++
	(*q.mux).Unlock()
}

//IsEmpty the Queue?
func (q *Queue) IsEmpty() bool {
	if q.Size == 0 {
		return true
	}
	return false
}

//IsMax have the minimum data in the queue?
func (q *Queue) IsMax() bool {
	if q.Size >= 500 {
		return true
	}
	return false
}

//DropRange have the minimum data in the queue?
func (q *Queue) DropRange(ran int) bool {

	if q.Size == 0 {
		return true
	}
	(*q.mux).Lock()
	if q.Size > ran {
		q.data = q.data[ran-1:]
	} else {
		return true
	}
	q.Size = q.Size - ran
	(*q.mux).Unlock()
	return true
}

func (q *Queue) GetData() []interface{} {
	return q.data
}
