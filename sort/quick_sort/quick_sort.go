package quick_sort

func Sort(itemList []int) []int {
	return quickSort(itemList, 0, len(itemList)-1)
}

func quickSort(itemList []int, low int, high int) []int {
	if low < high {
		itemList, pivotIndex := partition(itemList, low, high)

		return quickSort(
			quickSort(itemList, 0, pivotIndex-1),
			pivotIndex+1,
			high,
		)
	}

	return itemList
}

func partition(itemList []int, low int, high int) ([]int, int) {
	pivot := itemList[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if itemList[j] < pivot {
			i++

			itemList = swap(itemList, i, j)
		}
	}

	i++
	itemList = swap(itemList, i, high)

	return itemList, i
}

func swap(itemList []int, i int, j int) []int {
	if i == j {
		return itemList
	}

	tempValue := itemList[i]

	itemList[i] = itemList[j]
	itemList[j] = tempValue

	return itemList
}
