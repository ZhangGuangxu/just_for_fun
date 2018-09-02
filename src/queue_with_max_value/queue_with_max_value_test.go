package zgxutil

import (
    "testing"
)

func TestQueueWithMaxValue(t *testing.T) {
    q := NewQueueWithMaxValue()
    q.EnQueue(1)
    q.dump()
    q.EnQueue(2)
    q.dump()
    q.EnQueue(10)
    q.dump()
    q.EnQueue(9)
    q.dump()
    q.EnQueue(7)
    q.dump()
    q.EnQueue(10)
    q.dump()
    q.EnQueue(11)
    q.dump()
    q.EnQueue(3)
    q.dump()
    q.EnQueue(2)
    q.dump()
    q.EnQueue(8)
    q.dump()
    q.EnQueue(15)
    q.dump()
    q.EnQueue(6)
    q.dump()

    a, _ := q.MaxElement()
    if a != 15 {
        t.Error("should be 15")
    }

    q.DeQueue()
    a, _ = q.MaxElement()
    if a != 15 {
        t.Error("should be 15")
    }

    q.DeQueue()
    q.DeQueue()
    q.DeQueue()
    q.DeQueue()
    q.DeQueue()
    q.DeQueue()
    q.DeQueue()
    q.DeQueue()
    q.DeQueue()
    q.DeQueue()
    a, _ = q.MaxElement()
    if a != 6 {
        t.Error("should be 6")
    }

    q.EnQueue(10)
    a, _ = q.MaxElement()
    if a != 10 {
        t.Error("should be 10")
    }
}
