package containsduplicateii

func ContainsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)

	for i, n := range nums {
		if _, exist := numMap[n]; !exist {
			numMap[n] = i
			continue
		}

		if i-numMap[n] <= k {
			return true
		}
		numMap[n] = i
	}

	return false
}
