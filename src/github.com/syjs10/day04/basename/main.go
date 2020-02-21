package main

import "fmt"

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
func basename1(s string) {
	slast := strings.lastIndex(s, '\')
	s = s[slast+1:]
	if dot := strings.lastIndex(s, '.'); dot >= 0 {
		s = s[:dot]
	}
	return s
}
func main() {
	s := "/c.go"
	fmt.Println(basename(s))
	fmt.Println(basename1(s))
}
