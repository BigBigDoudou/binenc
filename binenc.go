// Package binenc encodes and decodes strings to integer basing on a fixed strings key.
package binenc

import "errors"

// Key is used to encode and decode elements (relying on indexes).
// Key should contain a maximum of 64 strings.
type Key []string

// Encode strings into int64 basing on the key.
// Elements not containing on the key are ignored.
// Return an error if element index is greater than 63.
// Elements order is ignored: ["foo", "bar"] and ["bar", "foo"] have the same code.
func (key Key) Encode(elements []string) (int64, error) {
	var res int64
	// for each element
	for _, element := range elements {
		// find the index of the element in the key
		for i, v := range key {
			if element == v {
				if i > 63 {
					return 0, errors.New("Index to high, key should contain 64 strings or less")
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

// Decode strings from int64 basing on the key.
// Strings are returned in the order as they appear in the key.
func (key Key) Decode(code int64) []string {
	res := []string{}
	// for each item, check if the related binary is 1
	// example:
	// 	if item index (i) == 2
	// 	check if the code third number from the right is 1:
	//  remove i numbers on the right then check if last numer is 1 (%2 == 1)
	//  examples
	// 		when code == 0101: 0101 >> 2 -> 01 -> add item
	// 		when code == 1001: 1001 >> 2 -> 10 -> continue
	for i, v := range key {
		if ((code >> i) % 2) == 1 {
			res = append(res, v)
		}
	}
	return res
}
