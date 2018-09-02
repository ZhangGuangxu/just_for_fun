package zgxutil

import (
    "testing"
)

func TestStackWithMaxValueTrace(t *testing.T) {
    {
        s := NewStackWithMaxValueTrace()
        s.Push(1)
        s.Push(2)
        s.Push(3)
        s.Push(4)
        s.Push(5)
        m, _ := s.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }

    {
        s := NewStackWithMaxValueTrace()
        s.Push(5)
        s.Push(4)
        s.Push(3)
        s.Push(2)
        s.Push(1)
        m, _ := s.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }

    {
        s := NewStackWithMaxValueTrace()
        s.Push(3)
        s.Push(4)
        s.Push(5)
        s.Push(2)
        s.Push(1)
        m, _ := s.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }

    {
        s := NewStackWithMaxValueTrace()
        s.Push(3)
        s.Push(4)
        s.Push(5)
        s.Push(2)
        s.Push(1)
        m, _ := s.Max()
        if m != 5 {
            t.Error("should be 5")
        }

        s.Push(10)
        m, _ = s.Max()
        if m != 10 {
            t.Error("should be 10")
        }
    }

    {
        s := NewStackWithMaxValueTrace()
        s.Push(1)
        s.Push(2)
        s.Push(3)
        s.Push(4)
        s.Push(5)
        m, _ := s.Max()
        if m != 5 {
            t.Error("should be 5")
        }

        s.Pop()
        m, _ = s.Max()
        if m != 4 {
            t.Error("should be 4")
        }
    }

    {
        s := NewStackWithMaxValueTrace()
        s.Push(3)
        s.Push(4)
        s.Push(5)
        s.Push(2)
        s.Push(1)
        m, _ := s.Max()
        if m != 5 {
            t.Error("should be 5")
        }

        s.Pop()
        m, _ = s.Max()
        if m != 5 {
            t.Error("should be 5")
        }
    }
}
