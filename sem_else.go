package main

import (
	"errors"
	"log"
)

func SemElse(numeric int64) int64 {
	if numeric%2 != 0 {

		var err = errorDivide()

		if err != nil {
			log.Println(err)
		}

	}

	return numeric
}

func errorDivide() error {
	return errors.New("error")
}
