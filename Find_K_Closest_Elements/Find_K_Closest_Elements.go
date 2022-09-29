package findkclosestelements

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func FindClosestElements(arr []int, k int, x int) []int {
	result := make([]int, 0, k)
	left, right := 0, len(arr)-1

	for right-left >= k {
		if abs(x-arr[left]) > abs(x-arr[right]) {
			left++
		} else {
			right--
		}
	}

	for i := left; i <= right; i++ {
		result = append(result, arr[i])
	}

	return result
}
