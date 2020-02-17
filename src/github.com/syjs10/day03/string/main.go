package main

import "fmt"

func main() {
	fmt.Println("hello\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("\u4e16\u754c")
	fmt.Println(HasPrefix("hello\xe4\xb8\x96\xe7\x95\x8c", "hello"))
	fmt.Println(HasSuffix("hello\xe4\xb8\x96\xe7\x95\x8c", "\u4e16\u754c"))
	fmt.Println(Contains("123hello\xe4\xb8\x96\xe7\x95\x8c312", "\u4e16\u754c"))
}

/*
 * 判断字符串是否包含前缀
 */
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

/*
 * 判断字符串是否包含后缀
 */
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

/*
 * 判断字符串是否包含子串
 */
func Contains(s, subStr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], subStr) {
			return true
		}
	}
	return false
}
