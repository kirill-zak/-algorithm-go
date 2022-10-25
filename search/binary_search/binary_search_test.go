package binary_search

import (
	"math/rand"
	"sort"
	"testing"
)

const NOT_FOUND_INDEX_VALUE = -1

type BinarySearchCase struct {
	ItemList      []int
	Searched      int
	ExpectedValue int
}

func TestBinarySearch(t *testing.T) {
	testCasesItemListSize := [4]int{45, 68, 110, 190}

	var caseList []BinarySearchCase

	for _, size := range testCasesItemListSize {
		caseList = append(caseList, makeItemList(size))
	}

	for _, testCase := range caseList {
		result, findFlag := BinarySearch(testCase.ItemList, testCase.Searched)

		if findFlag != true {
			t.Errorf("value [%d] not found", testCase.Searched)
		}

		if result != testCase.ExpectedValue {
			t.Errorf("value should be [%d], but got [%d]", testCase.ExpectedValue, result)
		}
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	testCasesItemListSize := [2]int{67, 126}

	var caseList []BinarySearchCase

	for _, size := range testCasesItemListSize {
		caseList = append(caseList, makeItemListNorFound(size))
	}

	for _, testCase := range caseList {
		result, findFlag := BinarySearch(testCase.ItemList, testCase.Searched)

		if findFlag != false {
			t.Errorf("value [%d] found", testCase.Searched)
		}

		if result != testCase.ExpectedValue {
			t.Errorf("value should be [%d], but got [%d]", testCase.ExpectedValue, result)
		}
	}
}

func makeItemList(itemListSize int) BinarySearchCase {
	var itemList []int

	for i := 0; i < itemListSize; i++ {
		itemList = append(itemList, rand.Int())
	}

	sort.Ints(itemList)

	searchedIndex := rand.Intn(itemListSize)

	return BinarySearchCase{
		ItemList:      itemList,
		Searched:      itemList[searchedIndex],
		ExpectedValue: searchedIndex,
	}
}

func makeItemListNorFound(itemListSize int) BinarySearchCase {
	var itemList []int

	for i := 0; i < itemListSize; i++ {
		itemList = append(itemList, rand.Int())
	}

	sort.Ints(itemList)

	searchedIndex := rand.Intn(itemListSize)

	return BinarySearchCase{
		ItemList:      itemList,
		Searched:      itemList[searchedIndex] * -1,
		ExpectedValue: NOT_FOUND_INDEX_VALUE,
	}
}
