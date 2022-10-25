package linear_search

const NotFoundIndex = -1

func Search(itemList []int, searchedItem int) (int, bool) {
	for index, value := range itemList {
		if value == searchedItem {
			return index, true
		}
	}

	return NotFoundIndex, false
}
