package main

// 给定一个字符串，如acabbdeff。找到字符串中只出现一次，并且是第一次出现的字符。

import (
	"errors"
	"fmt"
)

func main() {
	s := "acabbdeff"
	t, err := findFirst(s)
	fmt.Printf("%c, %v\n", t, err)

	s = "aabbccddeef"
	t, err = findFirst(s)
	fmt.Printf("%c, %v\n", t, err)

	s = "zaabbccddeef"
	t, err = findFirst(s)
	fmt.Printf("%c, %v\n", t, err)

	s = ""
	t, err = findFirst(s)
	fmt.Printf("%c, %v\n", t, err)

	s = "aabbccddee"
	t, err = findFirst(s)
	fmt.Printf("%c, %v\n", t, err)
}

var errEmptyString = errors.New("empty string")
var errNotFound = errors.New("not found")

func findFirst(s string) (byte, error) {
	if len(s) == 0 {
		return 0, errEmptyString
	}

	// 记录每个字符出现次数
	counts := make(map[byte]int)

	// 记录字符第一次出现的顺序
	seq := make([]byte, 0)

	for _, c := range s {
		b := byte(c)
		if cnt, ok := counts[b]; ok {
			if cnt == 1 {
				counts[b]++
			}
		} else {
			counts[b] = 1
			seq = append(seq, b)
		}
	}

	for _, b := range seq {
		if cnt, ok := counts[b]; ok && cnt == 1 {
			return b, nil
		}
	}

	return 0, errNotFound
}
