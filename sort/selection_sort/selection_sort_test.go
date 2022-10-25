package selection_sort

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

type TestCase struct {
	name string
	args args
	want []int
}

func TestSort(t *testing.T) {
	var tests []TestCase

	for _, size := range [4]int{65, 48, 369, 1257} {
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

func makeTestCase(size int) TestCase {
	var itemList []int

	for i := 0; i < size; i++ {
		itemList = append(itemList, rand.Int())
	}

	resultItemList := make([]int, len(itemList))

	copy(resultItemList, itemList)
	sort.Ints(resultItemList)

	return TestCase{
		name: fmt.Sprintf("test case with [%d] items", size),
		args: args{
			itemList: itemList,
		},
		want: resultItemList,
	}
}
