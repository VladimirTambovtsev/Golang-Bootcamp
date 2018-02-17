package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Jake", "Dan", "Alex", "John"}

	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
