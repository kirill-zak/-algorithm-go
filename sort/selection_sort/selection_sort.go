package selection_sort

func Sort(itemList []int) []int {
	for i := range itemList {
		minimumIndex := i

		for j := minimumIndex + 1; j < len(itemList); j++ {
			if itemList[minimumIndex] > itemList[j] {
				minimumIndex = j
			}
		}

		if minimumIndex != i {
			minimumValue := itemList[minimumIndex]
			itemList[minimumIndex] = itemList[i]
			itemList[i] = minimumValue
		}
	}

	return itemList
}
