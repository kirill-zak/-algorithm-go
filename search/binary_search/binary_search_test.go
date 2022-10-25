package binary_search

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

type args struct {
	list         []int
	searchedItem int
}

type TestCase struct {
	name         string
	args         args
	expected     int
	expectedFlag bool
}

func TestSearch(t *testing.T) {
	var tests []TestCase

	for _, size := range [4]int{549, 557, 3, 456} {
		tests = append(tests, makeTestCase(size))
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotResultFlag := Search(tt.args.list, tt.args.searchedItem)
			if gotResult != tt.expected {
				t.Errorf("Search() gotResult = %v, expected %v", gotResult, tt.expected)
			}
			if gotResultFlag != tt.expectedFlag {
				t.Errorf("Search() gotResultFlag = %v, expected %v", gotResultFlag, tt.expectedFlag)
			}
		})
	}
}

func makeTestCase(size int) TestCase {
	var itemList []int

	for i := 0; i < size; i++ {
		itemList = append(itemList, rand.Int())
	}

	sort.Ints(itemList)

	searchedIndex := rand.Intn(size)

	return TestCase{
		name: fmt.Sprintf("test case with [%d] items", size),
		args: args{
			list:         itemList,
			searchedItem: itemList[searchedIndex],
		},
		expected:     searchedIndex,
		expectedFlag: true,
	}
}

func TestSearchNotFound(t *testing.T) {
	var tests []TestCase

	for _, size := range [4]int{64, 957, 236, 6947} {
		tests = append(tests, makeTestCaseNotFound(size))
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotExpectedFlag := Search(tt.args.list, tt.args.searchedItem)
			if gotResult != tt.expected {
				t.Errorf("Search() gotResult = %v, want %v", gotResult, tt.expected)
			}
			if gotExpectedFlag != tt.expectedFlag {
				t.Errorf("Search() gotExpectedFlag = %v, want %v", gotExpectedFlag, tt.expectedFlag)
			}
		})
	}
}

func makeTestCaseNotFound(size int) TestCase {
	var itemList []int

	for i := 0; i < size; i++ {
		itemList = append(itemList, rand.Int())
	}

	sort.Ints(itemList)

	return TestCase{
		name: fmt.Sprintf("not found test case with [%d] items", size),
		args: args{
			list:         itemList,
			searchedItem: rand.Int() * -1,
		},
		expected:     NotFoundIndex,
		expectedFlag: false,
	}
}
