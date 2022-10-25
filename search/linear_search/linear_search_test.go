package linear_search

import (
	"fmt"
	"math/rand"
	"testing"
)

type args struct {
	itemList     []int
	searchedItem int
}

type TestCase struct {
	name       string
	args       args
	expected   int
	resultFlag bool
}

func TestSearch(t *testing.T) {
	var tests []TestCase

	for _, size := range [5]int{48, 26, 74, 18, 415} {
		tests = append(tests, makeTestCase(size))
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotResultFlag := Search(tt.args.itemList, tt.args.searchedItem)
			if gotResult != tt.expected {
				t.Errorf("Search() gotResult = %v, expected %v", gotResult, tt.expected)
			}
			if gotResultFlag != tt.resultFlag {
				t.Errorf("Search() gotResultFlag = %v, expected %v", gotResultFlag, tt.resultFlag)
			}
		})
	}
}

func makeTestCase(size int) TestCase {
	var itemList []int

	for i := 0; i < size; i++ {
		itemList = append(itemList, rand.Int())
	}

	searchedIndex := rand.Intn(size)

	return TestCase{
		name: fmt.Sprintf("test case with [%d] items", size),
		args: args{
			itemList:     itemList,
			searchedItem: itemList[searchedIndex],
		},
		expected:   searchedIndex,
		resultFlag: true,
	}
}

func TestSearchNotFound(t *testing.T) {
	var tests []TestCase

	for _, size := range [5]int{57, 12, 94, 578, 326} {
		tests = append(tests, makeTestCaseNotFound(size))
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotFlag := Search(tt.args.itemList, tt.args.searchedItem)
			if gotResult != tt.expected {
				t.Errorf("Search() gotResult = %v, want %v", gotResult, tt.expected)
			}
			if gotFlag != tt.resultFlag {
				t.Errorf("Search() gotFlag = %v, want %v", gotFlag, tt.resultFlag)
			}
		})
	}
}

func makeTestCaseNotFound(size int) TestCase {
	var itemList []int

	for i := 0; i < size; i++ {
		itemList = append(itemList, rand.Int())
	}

	return TestCase{
		name: fmt.Sprintf("not found test case with [%d] items", size),
		args: args{
			itemList:     itemList,
			searchedItem: rand.Int() * -1,
		},
		expected:   NotFoundIndex,
		resultFlag: false,
	}
}
