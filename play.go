package main

import "fmt"

type test struct {
	A int
	B int
}

func main() {
	mp := make(map[int]*test)
	mp[1] = &test{A: 0, B: 0}
	for i := 0; i < 11; i++ {
		mp[1].A += i
	}
	fmt.Println(*mp[1])
	for i := 0; i < 11; i++ {
		mp[1].B++
	}
	fmt.Println(*mp[1])
}
