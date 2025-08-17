package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

func process() func(a int) int {
	return func(a int) int {
		return a
	}
}
func main() {
	// result := add(4, 5)
	// fmt.Println(result)
	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)
	fn := process()
	fmt.Println(fn(6))
}
