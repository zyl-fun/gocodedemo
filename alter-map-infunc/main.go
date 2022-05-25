package main

import "fmt"

func main() {
	m := map[int]string{
		1: "1",
		2: "2",
	}
	fmt.Println("alter before:", m)

	altermap := func(m map[int]string) {
		m[1] = "1000"

	}
	altermap(m)
	fmt.Println("alter after: ", m)

}

/*
alter before: map[1:1 2:2]
alter after:  map[1:1000 2:2]
*/
