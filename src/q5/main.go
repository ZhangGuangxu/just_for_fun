package main

import (
	"errors"
	"fmt"
	"github.com/ZhangGuangxu/gds/src/slist"
)

func main() {
	{
		var s *slist.Slist
		SlistSelectionSort(s)
	}

	{
		s := slist.NewSlist()
		s.PushHead(1)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}

	{
		s := slist.NewSlist()
		s.PushHead(1)
		s.PushHead(2)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}

	{
		s := slist.NewSlist()
		s.PushHead(1)
		s.PushHead(2)
		s.PushHead(3)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}

	{
		s := slist.NewSlist()
		for i := 1; i < 10; i++ {
			s.PushHead(i)
		}
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}

	{
		s := slist.NewSlist()
		s.PushHead(1)
		s.PushHead(10)
		s.PushHead(20)
		s.PushHead(15)
		s.PushHead(16)
		s.PushHead(12)
		s.PushHead(2)
		s.PushHead(6)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	{
		var s *slist.Slist
		SlistSelectionSort2(s)
	}

	{
		s := slist.NewSlist()
		s.PushHead(1)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort2(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}

	{
		s := slist.NewSlist()
		s.PushHead(1)
		s.PushHead(2)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort2(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}

	{
		s := slist.NewSlist()
		s.PushHead(1)
		s.PushHead(2)
		s.PushHead(3)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort2(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}

	{
		s := slist.NewSlist()
		for i := 1; i < 10; i++ {
			s.PushHead(i)
		}
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort2(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}

	{
		s := slist.NewSlist()
		s.PushHead(1)
		s.PushHead(10)
		s.PushHead(20)
		s.PushHead(15)
		s.PushHead(16)
		s.PushHead(12)
		s.PushHead(2)
		s.PushHead(6)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()

		SlistSelectionSort2(s)
		for p := s.Head; p != nil; p = p.Next {
			fmt.Printf("%v ", p.Data)
		}
		fmt.Println()
	}
}

var errType = errors.New("type error")

// SlistSelectionSort sorts Slist with simple selection sort.
// Just swap data of items.
func SlistSelectionSort(s *slist.Slist) error {
	if s == nil {
		return nil
	}

	var end *slist.Item
	for s.Head != end {
		maxP := s.Head
		maxV, ok := s.Head.Data.(int)
		if !ok {
			return errType
		}
		prev := s.Head
		for i := s.Head.Next; i != end; i = i.Next {
			d, ok := i.Data.(int)
			if !ok {
				return errType
			}
			if d > maxV {
				maxP = i
				maxV = d
			}
			prev = i
		}
		if maxP != prev {
			maxP.Data, prev.Data = prev.Data, maxV
		}
		end = prev
	}

	return nil
}

// SlistSelectionSort2 sorts Slist with simple selection sort.
// Move items.
func SlistSelectionSort2(s *slist.Slist) error {
	if s == nil {
		return nil
	}

	var end *slist.Item
	for s.Head != end {
		var maxPrev *slist.Item // 记录当前最大值节点的前一个节点
		maxP := s.Head
		maxV, ok := s.Head.Data.(int)
		if !ok {
			return errType
		}
		prev := s.Head
		for i := s.Head.Next; i != end; i = i.Next {
			d, ok := i.Data.(int)
			if !ok {
				return errType
			}
			if d > maxV {
				maxPrev = prev
				maxP = i
				maxV = d
			}
			prev = i
		}
		if maxP != prev {
			// 先把maxP移除
			if maxPrev == nil {
				s.Head = maxP.Next
			} else {
				maxPrev.Next = maxP.Next
			}

			// 再连接到以prev结束的子链表尾部
			prev.Next = maxP
			maxP.Next = end
		}
		end = maxP
	}

	return nil
}
