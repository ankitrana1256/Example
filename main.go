package main

// Recursive function to find factorial
func factorial(x int) int {
	if x == 1 {
		return 1
	}

	return x * factorial(x-1)
}

// Overlapping String
func CommonStr(a, b string) string {
	var r string
	for i := 0; i < len(a); i++ {
		var s string
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				s += string(a[i])
			} else {
				break
			}
		}
		if len(s) > len(r) {
			r = s
		}
	}
	return r
}

func main() {
	CommonStr("Hello", "He")
}
