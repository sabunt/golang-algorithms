package two_sum

import "sort"
//If we return indexes of array


// first my solution
//Time complexity : O(n^2)
//Space complexity : O(n).
func twoSum1(nums []int, target int) []int {

	m := map[int][]int{}
	for index1, element1 := range nums {
		nextIndex := index1 + 1
		if index1 == len(nums) {
		 break
		}
		for index2, element2 := range nums[nextIndex:] {

			sum := element1 + element2
			if _, ok := m[sum]; ok {
				continue
			}

			m[sum] = []int{index1, nextIndex+index2}
		}
	}
	if m[target] == nil {
			return []int{0,0}
			}

	return m[target]
}

//Approach 1: Brute Force
//The brute force approach is simple. Loop through each element xx and find if there is another value that equals to target - x.
//
//Time complexity : O(n^2)
//For each element, we try to find its complement by looping through the rest of array which takes O(n) time.
//Therefore, the time complexity is O(n^2)
//
//Space complexity : O(1)


func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for i2 := i + 1; i2 < len(nums); i2++ {
			if nums[i2] == target - nums[i] {
				return []int{ i, i2 }

			}
		}
	}
	return []int{0,0}
}

// O(n) time, O(n) space
func twoSum3(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := m[complement]; ok {
			return []int{m[complement],i}
		}
		m[nums[i]] = i //save result for checking in next iteration
	}
	return []int{0,0}
}



//If we return elements of array

//TwoNumberSum1 O(n^2) time complexity, O(1) space complexity
func TwoNumberSum1(array []int, target int) []int {
	for i, value := range array {
		for _, value2 := range array[i+1:] {
			complement := target - value
			if complement == value2{
				return []int{value,value2}
			}
		}
	}
	return []int{}
}



//TwoNumberSum2 O(n) time, O(n) space
func TwoNumberSum2(array []int, target int) []int {
	m := map[int]bool{}
	for _, value := range array {
		complement := target - value
		if _ , ok := m[complement]; ok {
			return []int{complement, value}
		}
		m[value] = true
	}
	return []int{}
}

//TwoNumberSum3 O(n log (n)) time, O(1) space
func TwoNumberSum3(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		sum := array[left] + array[right]
		if sum == target {
			return []int{array[left], array[right]}
		} else if sum < target{
			left ++
		} else{
		 right --
		}

	}
	return []int{}
} 

