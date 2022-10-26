# How to use

## Import

```go
QuickSort "github.com/kirill-zak/algorithm-go/sort/quick_sort"
```

## Get result

```go
itemList = QuickSort.Sort(itemList)
```

### Example

```go
package main

import (
	"fmt"
	"math/rand"

	QuickSort "github.com/kirill-zak/algorithm-go/sort/quick_sort"
)

const Size = 6

func main() {
	var itemList []int

	for i := 0; i < Size; i++ {
		itemList = append(itemList, rand.Int())
	}

	fmt.Println("Before sort")
	for _, value := range itemList {
		println("* %d", value)
	}

	itemList = QuickSort.Sort(itemList)

	fmt.Println("After sort")
	for _, value := range itemList {
		println("* %d", value)
	}
}
```