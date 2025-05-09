package main

import "fmt"

type field struct {
	num int
}

func (t *field) print(n int) {
	fmt.Println(t.num, n)
}
func main() {
	var i int = 1
	defer fmt.Println("result2 =>", func() int { return i * 2 }())
	i++

	v := field{1}
	defer v.print(func() int { return i * 2 }())
	v = field{2}
	i++

	// prints:
	// 2 4
	// result => 2 (not ok if you expected 4)
}
