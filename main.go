package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		var result = SemElse(int64(i))
		fmt.Println(result)
	}
}
