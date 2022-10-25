# How to use

## Import

```go
import BinarySearch "github.com/kirill-zak/algorithm-go/search/binary_search"
```

## Get result

```go
_, foundFlag := BinarySearch.Search(itemList, 42)
```

### Example

```go
package main

import (
	"fmt"
	BinarySearch "github.com/kirill-zak/algorithm-go/search/binary_search"
)

const SearchedItem = 23

func main() {
	var itemList []int

	for i := 0; i < 45; i++ {
		itemList = append(itemList, i)
	}

	result, foundFlag := BinarySearch.Search(itemList, SearchedItem)
	if foundFlag != true {
		fmt.Printf("item [%d] not found", SearchedItem)
	} else {
		fmt.Printf("item [%d] found with position [%d]", SearchedItem, result)
	}
}
```