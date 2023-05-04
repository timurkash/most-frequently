package main

import (
	"fmt"
	"sort"
)

var (
	sl = []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 5}
	l  = 2
)

func main() {
	m := make(map[int]int)
	var ok bool
	for _, s := range sl {
		if _, ok = m[s]; ok {
			m[s]++
		} else {
			m[s] = 1
		}
	}
	m2 := make(map[int][]int)
	var res2 []int
	for k, v := range m {
		if _, ok = m2[v]; !ok {
			m2[v] = []int{k}
			res2 = append(res2, v)
		} else {
			m2[v] = append(m2[v], k)
		}
	}
	sort.Ints(res2)
	lr2 := len(res2)
	if l > lr2 {
		l = lr2
	}
	var res []int
	for i := 0; i < l; i++ {
		res = append(res, m2[res2[lr2-i-1]]...)
	}
	fmt.Println(res)
}
