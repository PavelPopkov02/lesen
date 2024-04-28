package main

import "fmt"

func main() {
	var r int
	fmt.Scan(&r)
	h := r / 30
	m := (r % 30) * 2
	fmt.Println("It is", h, "hours", m, "minets.")
}
