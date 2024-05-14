package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	//
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			fmt.Println(j)
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
