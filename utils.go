package main

func isLetterInWord(l string, w string) bool {
	for _, ch := range w {
		if string(ch) == l {
			return true
		}
	}

	return false
}
