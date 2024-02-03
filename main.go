package main

import (
	"fmt"
	"log"
)

func main() {
	for i := 0; i < 10; i++ {
		var result, err = WithElse(int64(i))

		if err != nil {
			log.Println(err)
		}

		fmt.Println(result)
	}
}
