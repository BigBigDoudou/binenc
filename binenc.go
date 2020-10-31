// Package binenc encodes and decodes strings to integer basing on a fixed strings list.
package binenc

import "errors"

// List is used to encode and decode elements where elements are a slice of the list.
// Enconding and decoding rely on indexes so changing items order would break the code system.
// List should contain a maximum of 64 strings.
type List []string

// Encode encodes strings to integer basing on the list.
// Elements not containing on the list will be ignored.
// Return an error if element index is the list is greater than 63.
// Elements order is not encoded: ["foo", "bar"] and ["bar", "foo"] have the same code.
func (list List) Encode(elements []string) (int64, error) {
	var res int64
	// for each element
	for _, element := range elements {
		// find the index of the element in the list array
		for i, item := range list {
			if element == item {
				if i > 63 {
					return 0, errors.New("Index to high, list should contain 64 strings or less")
				}
				// add to res 1 left shifted by the index
				// examples:
				// 	 when i == 0: 1 << 0 -> 1   -> res += 1 (1 base2   -> 1 base10)
				// 	 when i == 2: 1 << 2 -> 100 -> res += 4 (100 base2 -> 4 base10)
				res += 1 << i
				break
			}
		}
	}
	return res, nil
}

// Decode decodes strings from integer basing on the list.
// Strings are returned in the order as they appear in the list.
func (list List) Decode(code int64) []string {
	res := []string{}
	// for each item, check if the related binary is 1
	// example:
	// 	if item index (i) == 2
	// 	check if the code third number from the right is 1:
	//  remove i numbers on the right then check if last numer is 1 (%2 == 1)
	//  examples
	// 		when code == 0101: 0101 >> 2 -> 01 -> add item
	// 		when code == 1001: 1001 >> 2 -> 10 -> continue
	for i, item := range list {
		if ((code >> i) % 2) == 1 {
			res = append(res, item)
		}
	}
	return res
}
