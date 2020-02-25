package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["js"] = 25
	ages["av"] = 23
	ages["vc"] = 24
	ages["hg"] = 27
	var names []string
	// for name, age := range ages {
	// 	fmt.Printf("%s\t%d\n", name, age)
	// }
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
