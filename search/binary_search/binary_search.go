package binary_search

const NotFoundIndex = -1

func BinarySearch(list []int, searchedItem int) (int, bool) {
	low := 0
	high := len(list) - 1

	for low <= high {
		middle := (low + high) / 2

		middleValue := list[middle]

		if middleValue == searchedItem {
			return middle, true
		}

		if middleValue > searchedItem {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return NotFoundIndex, false
}
