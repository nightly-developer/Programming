package main

import ("fmt")

func findIndex(list []int,item int) int {
	for j := 0; j < len(list); j++ {
		if list[j] == item {
			return j
		}
	}
	return -1
}

func twoSum(nums[]int,target int) []int {
	var numMap = map[int]int{}
	// var indices []int = []int{}

	for i, num := range nums{
		var complement int = target - num
		if index,ok := numMap[complement];ok{
			return []int{index,i}
		}
		numMap[num] = i
	}
	return []int{}
}

func main(){
	var nums []int = []int{3,2,4}
	const target = 6
	fmt.Println(twoSum(nums,target))
}