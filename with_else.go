package main

import "errors"

func WithElse(numeric int64) (int64, error) {
	if numeric%2 == 0 {
		return numeric, nil
	} else {
		return 0, errors.New("error ao dividir")
	}
}
