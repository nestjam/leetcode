package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		j, ok := m[nums[i]]
		if ok {
			return []int{j, i}
		}
		m[target-nums[i]] = i
	}
	panic("oops, invalid nums")
}
