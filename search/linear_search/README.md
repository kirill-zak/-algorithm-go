# How to use

## Import

```go
import LinearSearch "github.com/kirill-zak/algorithm-go/search/linear_search"
```

## Get result

```go
_, foundFlag := LinearSearch.Search(itemList, 23)
```

### Example

```go
package main

import (
	"fmt"
	"math/rand"

	LinearSearch "github.com/kirill-zak/algorithm-go/search/linear_search"
)

const SearchedItem = 23

func main() {
	var itemList []int

	for i := 0; i < 45; i++ {
		itemList = append(itemList, rand.Int())
	}

	result, foundFlag := LinearSearch.Search(itemList, SearchedItem)
	if foundFlag != true {
		fmt.Printf("item [%d] not found", SearchedItem)
	} else {
		fmt.Printf("item [%d] found with position [%d]", SearchedItem, result)
	}
}
```