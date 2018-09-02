package zgxutil

import (
    "errors"
    "fmt"
    //"log"
)

var ErrEmptyQueue = errors.New("empty queue")

type item struct {
    Data int
    Next *item
}

func newItem(x int) *item {
    return &item{
        Data: x,
    }
}

type queue struct {
    head *item
    tail *item
}

func newQueue() *queue {
    return &queue{}
}

func (q *queue) IsEmpty() bool {
    return q.head == nil 
}

func (q *queue) PushTail(x int) *item {
    i := newItem(x)
    if q.IsEmpty() {
        q.head = i
    } else {
        q.tail.Next = i
    }
    q.tail = i
    return i
}

func (q *queue) PopHead() (*item, error) {
    if q.IsEmpty() {
        return nil, ErrEmptyQueue
    }

    i := q.head
    q.head = q.head.Next
    return i, nil
}

type QueueWithMaxValue struct {
    values *queue
    maxValues []*item
}

func NewQueueWithMaxValue() *QueueWithMaxValue {
    return &QueueWithMaxValue{
        values: newQueue(),
        maxValues: []*item{},
    }
}

func (q *QueueWithMaxValue) addMaxValues(i *item) {
    // 如果有相同元素就替换.
    // 如果没有相同元素就插入，然后丢弃比它小的所有元素.

    high := len(q.maxValues)-1
    low := 0
    mid := 0

	for low <= high {
		mid = (low+high)/2
		if q.maxValues[mid].Data == i.Data {
            q.maxValues[mid] = i
            return
		} else if q.maxValues[mid].Data > i.Data {
			low = mid + 1
		} else { // q.maxValues[mid].Data < i.Data
            high = mid - 1
		}
    }
    
    //log.Println(mid)

    var j int
    if q.maxValues[mid].Data > i.Data {
        j = mid + 1
        if j > len(q.maxValues)-1 {
            q.maxValues = append(q.maxValues, i)
        }
    } else { // q.maxValues[mid].Data < i.Data
        j = mid
    }
    q.maxValues[j] = i
    q.maxValues = q.maxValues[:j+1]
}

func (q *QueueWithMaxValue) EnQueue(v int) {
    i := q.values.PushTail(v)
    if len(q.maxValues) == 0 {
        q.maxValues = append(q.maxValues, i)
    } else {
        v := i.Data
        if v >= q.maxValues[0].Data {
            q.maxValues = []*item{i}
        } else if v < q.maxValues[len(q.maxValues)-1].Data {
            q.maxValues = append(q.maxValues, i)
        } else if v == q.maxValues[len(q.maxValues)-1].Data {
            q.maxValues[len(q.maxValues)-1] = i
        } else {
            q.addMaxValues(i)
        }
    }
}

func (q *QueueWithMaxValue) DeQueue() (int, error) {
    i, err := q.values.PopHead()
    if err != nil {
        return 0, err
    }
    if i == q.maxValues[0] {
        if len(q.maxValues) > 1 {
            q.maxValues = q.maxValues[1:]
        } else {
            q.maxValues = []*item{}
        }
    }
    return i.Data, nil
}

func (q *QueueWithMaxValue) MaxElement() (int, error) {
    if len(q.maxValues) == 0 {
        return 0, ErrEmptyQueue
    }

    return q.maxValues[0].Data, nil
}

func (q *QueueWithMaxValue) dump() {
    for j := 0; j < len(q.maxValues); j++ {
        fmt.Printf("%d ", q.maxValues[j].Data)
    }
    fmt.Println("")
}
