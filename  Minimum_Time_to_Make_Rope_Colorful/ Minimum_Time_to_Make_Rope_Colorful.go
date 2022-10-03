package minimumtimetomakeropecolorful

func MinCost(colors string, neededTime []int) int {
	minNeedTime := 0
	index, compareIndex := 0, 1

	for index < len(colors) && compareIndex < len(colors) {
		if colors[index] == colors[compareIndex] {
			if neededTime[index] < neededTime[compareIndex] {
				minNeedTime += neededTime[index]
				index = compareIndex
				compareIndex++

			} else {
				minNeedTime += neededTime[compareIndex]
				compareIndex++
			}
		} else {
			index = compareIndex
			compareIndex++
		}
	}

	return minNeedTime
}
