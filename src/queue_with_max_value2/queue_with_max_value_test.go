package zgxutil

import (
    "testing"
)

func TestQueueWithMaxValue(t *testing.T) {
    {
        q := NewQueueWithMaxValue()
        q.PushTail(1)
        q.PushTail(2)
        q.PushTail(3)
        q.PushTail(4)
        q.PushTail(5)
        m, _ := q.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }

    {
        q := NewQueueWithMaxValue()
        q.PushTail(5)
        q.PushTail(4)
        q.PushTail(3)
        q.PushTail(2)
        q.PushTail(1)
        m, _ := q.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }

    {
        q := NewQueueWithMaxValue()
        q.PushTail(3)
        q.PushTail(4)
        q.PushTail(5)
        q.PushTail(2)
        q.PushTail(1)
        m, _ := q.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }

    {
        q := NewQueueWithMaxValue()
        q.PushTail(3)
        q.PushTail(4)
        q.PushTail(5)
        q.PushTail(2)
        q.PushTail(1)
        m, _ := q.Max()
        if m != 5 {
            t.Error("should be 5")
        }

        q.PushTail(10)
        m, _ = q.Max()
        if m != 10 {
            t.Error("should be 10")
        }
    }

    {
        q := NewQueueWithMaxValue()
        q.PushTail(1)
        q.PushTail(2)
        q.PushTail(3)
        q.PushTail(4)
        q.PushTail(5)
        m, _ := q.Max()
        if m != 5 {
            t.Error("should be 5")
        }

        q.PopHead()
        m, _ = q.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }

    {
        q := NewQueueWithMaxValue()
        q.PushTail(3)
        q.PushTail(4)
        q.PushTail(5)
        q.PushTail(2)
        q.PushTail(1)
        m, _ := q.Max()
        if m != 5 {
            t.Error("should be 5")
        }

        q.PopHead()
        m, _ = q.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }

    {
        q := NewQueueWithMaxValue()
        q.PushTail(3)
        q.PushTail(4)
        q.PushTail(5)
        q.PushTail(2)
        q.PushTail(1)
        m, _ := q.Max()
        if m != 5 {
            t.Error("should be 5")
        }

        q.PopHead()
        q.PopHead()
        m, _ = q.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }

    {
        q := NewQueueWithMaxValue()
        q.PushTail(3)
        q.PushTail(4)
        q.PushTail(5)
        q.PushTail(2)
        q.PushTail(1)
        m, _ := q.Max()
        if m != 5 {
            t.Error("should be 5")
        }

        q.PopHead()
        q.PopHead()
        q.PopHead()
        m, _ = q.Max()
        if m != 2 {
            t.Error("should be 2")
        }
    }
}
