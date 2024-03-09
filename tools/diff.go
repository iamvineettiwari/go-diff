package tools

func longgestCommonSubsequence(firstItem, secondItem []string, firstIdx, secondIdx int, cache map[int]map[int][]string) []string {
	if firstIdx >= len(firstItem) || secondIdx >= len(secondItem) {
		return []string{}
	}

	if _, isFirstIdxCalculated := cache[firstIdx]; isFirstIdxCalculated {
		if calculatedResult, isSecondIdxCalculated := cache[firstIdx][secondIdx]; isSecondIdxCalculated {
			return calculatedResult
		}
	} else {
		cache[firstIdx] = make(map[int][]string)
	}

	if firstItem[firstIdx] == secondItem[secondIdx] {
		dataCollectedByMovingBothPointer := longgestCommonSubsequence(firstItem, secondItem, firstIdx+1, secondIdx+1, cache)
		dataCollectedByMovingBothPointer = append([]string{firstItem[firstIdx]}, dataCollectedByMovingBothPointer...)
		cache[firstIdx][secondIdx] = dataCollectedByMovingBothPointer
		return dataCollectedByMovingBothPointer
	}

	dataCollectedByMovingOnlyFirstPointer := longgestCommonSubsequence(firstItem, secondItem, firstIdx+1, secondIdx, cache)
	dataCollectedByMovingOnlySecondPointer := longgestCommonSubsequence(firstItem, secondItem, firstIdx, secondIdx+1, cache)

	firstPointerCollectedDataLen := len(dataCollectedByMovingOnlyFirstPointer)
	secondPointerCollectedDataLen := len(dataCollectedByMovingOnlySecondPointer)

	if firstPointerCollectedDataLen > secondPointerCollectedDataLen {
		cache[firstIdx][secondIdx] = dataCollectedByMovingOnlyFirstPointer
		return dataCollectedByMovingOnlyFirstPointer
	}

	cache[firstIdx][secondIdx] = dataCollectedByMovingOnlySecondPointer
	return dataCollectedByMovingOnlySecondPointer
}

func FindCommon(firstItem, secondItem []string) []string {
	cache := make(map[int]map[int][]string)
	return longgestCommonSubsequence(firstItem, secondItem, 0, 0, cache)
}

func FindDiff(firstItem, secondItem []string) []string {
	common := FindCommon(firstItem, secondItem)
	output := []string{}

	firstIdx, secondIdx := 0, 0

	for _, commonStr := range common {

		for firstIdx < len(firstItem) && firstItem[firstIdx] != commonStr {
			output = append(output, "< "+firstItem[firstIdx])
			firstIdx += 1
		}

		for secondIdx < len(secondItem) && secondItem[secondIdx] != commonStr {
			output = append(output, "> "+secondItem[secondIdx])
			secondIdx += 1
		}

		firstIdx += 1
		secondIdx += 1
	}

	for firstIdx < len(firstItem) {
		output = append(output, "< "+firstItem[firstIdx])
		firstIdx += 1
	}

	for secondIdx < len(secondItem) {
		output = append(output, "> "+secondItem[secondIdx])
		secondIdx += 1
	}

	return output
}
