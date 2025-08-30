package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", true
// }

// func process() func(a int) int {
// 	return func(a int) int {
// 		return a
// 	}
// }

func findDuplicates(nums []int) []int {
	res := make([]int, 0, len(nums))
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for key, val := range m {
		if val > 1 {
			res = append(res, key)
		}
	}
	return res
}
func main() {
	// result := add(4, 5)
	// fmt.Println(result)
	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)
	// fn := process()
	// fmt.Println(fn(6))
	nums := []int{1, 2, 2, 3, 4, 4, 5, 6, 6, 7}
	dup := findDuplicates(nums)
	fmt.Println(dup)

}
