# How to use

## Import

```go
import SelectionSort "github.com/kirill-zak/algorithm-go/sort/selection_sort"
```

## Get result

```go
itemList = SelectionSort.Sort(itemList)
```

### Example

```go
package main

import (
	"fmt"
	"math/rand"

	SelectionSort "github.com/kirill-zak/algorithm-go/sort/selection_sort"
)

const SIZE = 5

func main() {
	itemList := makeItemList()

	fmt.Println("got item list with values")
	for _, value := range itemList {
		fmt.Printf("* %d\n", value)
	}

	itemList = SelectionSort.Sort(itemList)

	fmt.Println("after sort")
	for _, value := range itemList {
		fmt.Printf("* %d\n", value)
	}
}

func makeItemList() []int {
	var itemList []int

	for i := 0; i < SIZE; i++ {
		itemList = append(itemList, rand.Int())
	}

	return itemList
}
```