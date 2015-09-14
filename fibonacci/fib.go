package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	limit := 1
	fmt.Scanf("%d", &limit)
	for i:=0; i<limit; i++ {
		fmt.Println(f())
	}
}
