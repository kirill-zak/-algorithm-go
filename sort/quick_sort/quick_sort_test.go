package quick_sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

type args struct {
	itemList []int
}

type testCase struct {
	name string
	args args
	want []int
}

func TestSort(t *testing.T) {
	var tests []testCase

	for _, size := range [1]int{5} {
		tests = append(tests, makeTestCase(size))
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.itemList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeTestCase(size int) testCase {
	var itemList []int

	for i := 0; i < size; i++ {
		itemList = append(itemList, rand.Int())
	}

	expectedValue := make([]int, len(itemList))
	copy(expectedValue, itemList)
	sort.Ints(expectedValue)

	return testCase{
		name: fmt.Sprintf("test case with [%d] items", size),
		args: args{
			itemList: itemList,
		},
		want: expectedValue,
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		itemList []int
		low      int
		high     int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case #1",
			args: args{
				itemList: []int{24, 68, 41, 98, 25},
				low:      0,
				high:     4,
			},
			want: []int{24, 25, 41, 68, 98},
		},
		{
			name: "Test case #2",
			args: args{
				itemList: []int{69, 88, 1, 74, 125, 66, 47},
				low:      0,
				high:     6,
			},
			want: []int{1, 47, 66, 69, 74, 88, 125},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.itemList, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		itemList []int
		low      int
		high     int
	}

	tests := []struct {
		name       string
		args       args
		want       []int
		pivotIndex int
	}{
		{
			name: "Test case #1",
			args: args{
				itemList: []int{10, 80, 30, 90, 40, 50, 70},
				low:      0,
				high:     6,
			},
			want:       []int{10, 30, 40, 50, 70, 90, 80},
			pivotIndex: 4,
		},
		{
			name: "Test case #2",
			args: args{
				itemList: []int{22, 14, 2, 34, 60, 110, 90, 111, 95, 80},
				low:      0,
				high:     9,
			},
			want:       []int{22, 14, 2, 34, 60, 80, 90, 111, 95, 110},
			pivotIndex: 5,
		},
		{
			name: "Test case #4 (sort right)",
			args: args{
				itemList: []int{22, 14, 2, 34, 60, 80, 90, 111, 95, 110},
				low:      6,
				high:     9,
			},
			want:       []int{22, 14, 2, 34, 60, 80, 90, 95, 110, 111},
			pivotIndex: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotIndex := partition(tt.args.itemList, tt.args.low, tt.args.high)
			if !reflect.DeepEqual(gotResult, tt.want) {
				t.Errorf("partition() gotResult = %v, want %v", gotResult, tt.want)
			}
			if gotIndex != tt.pivotIndex {
				t.Errorf("partition() gotIndex = %v, want %v", gotIndex, tt.pivotIndex)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		itemList []int
		i        int
		j        int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case #1",
			args: args{
				itemList: []int{3, 10, 46},
				i:        1,
				j:        2,
			},
			want: []int{3, 46, 10},
		},
		{
			name: "Test case #2",
			args: args{
				itemList: []int{1, 69, 47, 31, 58},
				i:        0,
				j:        3,
			},
			want: []int{31, 69, 47, 1, 58},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swap(tt.args.itemList, tt.args.i, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swap() = %v, want %v", got, tt.want)
			}
		})
	}
}
