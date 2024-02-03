package main

func SemElse(numeric int64) (int64, bool) {
	if numeric%2 == 0 {
		return numeric, false
	}

	return 0, true
}
