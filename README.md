# binenc

A binary encoder that encodes a slice of strings into a int64. This is useful if you need to persist a combination of options and can *not* afford to save an array in database.

The enconding bases on a `List` of strings which is the **key** to encode and decode. Changing the list would then obviously break the coding system. It should contain 64 strings or less. Providing a list with more than 64 strings returns a error on `Encode` call. The encoder generates a uniq code for each of the `18446744073709551615` combinations (for  64 strings) and can then decode it to return the original combination.

```go
package main

import (
	"binenc"
	"fmt"
	"strconv"
)

func main() {
	list := binenc.List{"alfa", "bravo", "charlie", "delta", "echo", "foxtrot", "golf"}

	encoded, _ := list.Encode([]string{"bravo", "delta", "foxtrot"})
	decoded := list.Decode(42)

	fmt.Println(encoded)                       // 42     (base 10)
	fmt.Println(strconv.FormatInt(encoded, 2)) // 101010 (base 2)
	fmt.Println(decoded)                       // [bravo delta foxtrot]
}
```