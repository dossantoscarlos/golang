package main

func WithElse(numeric int64) (int64, bool) {
	if numeric%2 == 0 {
		return numeric, false
	} else {
		return 0, true
	}
}
