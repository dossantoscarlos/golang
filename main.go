package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		var result, err = SemElse(int64(i))

		if !err {
			fmt.Println(result)
		}

	}
}
