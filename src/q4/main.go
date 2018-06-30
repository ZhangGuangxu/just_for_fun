package main

import (
	"fmt"
	"github.com/ZhangGuangxu/gds/src/slist"
)

// 给定单链表一个中间节点(不是最后一个节点)，但不知道单链表表头指针，删除那个中间节点。要求时间复杂度O(1)。

func main() {
	s := slist.NewSlist()
	for i := 0; i < 100; i++ {
		s.PushHead(i)
	}
	i := s.FindItem(51)
	if i == nil {
		fmt.Println("i should not be nil")
	}
	r := DelSlistNode(i)
	fmt.Println(r)
	i = s.FindItem(51)
	if i != nil {
		fmt.Println("i should be nil")
	}
}

// DelSlistNode deletes a node from the slist.
func DelSlistNode(i *slist.Item) bool {
	if i == nil {
		return false
	}

	n := i.Next
	if n == nil { // 如果i是单链表最后一个节点就无法处理
		return false
	}

	i.Data = n.Data
	i.Next = n.Next
	return true
}
