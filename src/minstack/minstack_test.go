package minstack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	{
		s := NewMinStack()
		s.Push(1)
		m, _ := s.Min()
		if m != 1 {
			t.Errorf("s.Min() got %d, want %d", m, 1)
		}
		s.Push(2)
		m, _ = s.Min()
		if m != 1 {
			t.Errorf("s.Min() got %d, want %d", m, 1)
		}
		s.Push(3)
		m, _ = s.Min()
		if m != 1 {
			t.Errorf("s.Min() got %d, want %d", m, 1)
		}
	}

	{
		s := NewMinStack()
		s.Push(10)
		m, _ := s.Min()
		if m != 10 {
			t.Errorf("s.Min() got %d, want %d", m, 10)
		}
		s.Push(9)
		m, _ = s.Min()
		if m != 9 {
			t.Errorf("s.Min() got %d, want %d", m, 9)
		}
		s.Push(8)
		m, _ = s.Min()
		if m != 8 {
			t.Errorf("s.Min() got %d, want %d", m, 8)
		}
	}

	{
		s := NewMinStack()
		s.Push(10)
		m, _ := s.Min()
		if m != 10 {
			t.Errorf("s.Min() got %d, want %d", m, 10)
		}
		s.Push(6)
		m, _ = s.Min()
		if m != 6 {
			t.Errorf("s.Min() got %d, want %d", m, 6)
		}
		s.Push(8)
		m, _ = s.Min()
		if m != 6 {
			t.Errorf("s.Min() got %d, want %d", m, 6)
		}
	}

	{
		s := NewMinStack()
		s.Push(10)
		s.Push(5)
		s.Push(7)
		s.Push(3)
		s.Push(4)
		s.Push(1)

		if v, _ := s.Min(); v != 1 {
			t.Errorf("s.Min() got %d, want %d", v, 1)
		}

		if v, _ := s.Pop(); v != 1 {
			t.Errorf("s.Pop() got %d, want %d", v, 1)
		}
		if v, _ := s.Min(); v != 3 {
			t.Errorf("s.Min() got %d, want %d", v, 3)
		}

		if v, _ := s.Pop(); v != 4 {
			t.Errorf("s.Pop() got %d, want %d", v, 4)
		}
		if v, _ := s.Min(); v != 3 {
			t.Errorf("s.Min() got %d, want %d", v, 3)
		}

		if v, _ := s.Pop(); v != 3 {
			t.Errorf("s.Pop() got %d, want %d", v, 3)
		}
		if v, _ := s.Min(); v != 5 {
			t.Errorf("s.Min() got %d, want %d", v, 5)
		}

		if v, _ := s.Pop(); v != 7 {
			t.Errorf("s.Pop() got %d, want %d", v, 7)
		}
		if v, _ := s.Min(); v != 5 {
			t.Errorf("s.Min() got %d, want %d", v, 5)
		}

		if v, _ := s.Pop(); v != 5 {
			t.Errorf("s.Pop() got %d, want %d", v, 5)
		}
		if v, _ := s.Min(); v != 10 {
			t.Errorf("s.Min() got %d, want %d", v, 10)
		}

		if v, _ := s.Pop(); v != 10 {
			t.Errorf("s.Pop() got %d, want %d", v, 10)
		}
		if _, err := s.Min(); err == nil {
			t.Error("s.Min() should return error")
		}
		if _, err := s.Pop(); err == nil {
			t.Error("s.Pop() should return error")
		}
	}
}

func BenchmarkMinStackPushPop(b *testing.B) {
	s := NewMinStack()
	s.Push(10)
	s.Push(5)
	s.Push(7)
	s.Push(3)
	s.Push(4)
	s.Push(1)

	for i := 0; i < b.N; i++ {
		s.Push(-1)
		s.Pop()
	}
}

func BenchmarkMinStackMin(b *testing.B) {
	s := NewMinStack()
	s.Push(10)
	s.Push(5)
	s.Push(7)
	s.Push(3)
	s.Push(4)
	s.Push(1)

	for i := 0; i < b.N; i++ {
		s.Min()
	}
}
