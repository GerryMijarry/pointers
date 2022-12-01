package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(i)

	p := &i

	fmt.Println(j)

	fmt.Println(*p)

	fmt.Printf("%T\n", p)
	*p = 21
	fmt.Println(i)
}
