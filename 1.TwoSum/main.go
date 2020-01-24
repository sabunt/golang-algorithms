package main
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


func twoSum3(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := m[complement]; ok { //сравниваем с прошлыми результатами
			return []int{m[complement],i}
		}
		m[nums[i]] = i //save result for checking in next iteration
	}
	return []int{0,0}
}