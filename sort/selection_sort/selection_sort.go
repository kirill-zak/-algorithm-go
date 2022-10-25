package selection_sort

func Sort(itemList []int) []int {
	for i, value := range itemList {
		minimumIndex := i

		for j := minimumIndex + 1; j < len(itemList); j++ {
			if itemList[minimumIndex] > itemList[j] {
				minimumIndex = j
			}
		}

		if minimumIndex != i {
			itemList[i] = itemList[minimumIndex]
			itemList[minimumIndex] = value
		}
	}

	return itemList
}
