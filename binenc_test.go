package binenc

import "testing"

func TestEncode(t *testing.T) {
	list := List{"alfa", "bravo", "charlie", "delta", "echo", "foxtrot", "golf"}
	example := []string{"bravo", "delta", "foxtrot"}
	want := int64(42)
	encoded, err := list.Encode(example)
	if want != encoded || err != nil {
		t.Fatalf(`Encode([]string{"bravo", "delta", "foxtrot"}) = %v, %v, want %v, nil`, encoded, err, want)
	}
}

func TestDecode(t *testing.T) {
	list := List{"alfa", "bravo", "charlie", "delta", "echo", "foxtrot", "golf"}
	example := int64(42)
	want := []string{"bravo", "delta", "foxtrot"}
	decoded := list.Decode(example)
	if !matchSlices(want, decoded) {
		t.Fatalf(`Decoded(int64(42)) = %v, want %v`, decoded, want)
	}
}

// Test if slices contains the same strings (not ordered)
func matchSlices(sliceX []string, sliceY []string) bool {
	if len(sliceX) != len(sliceY) {
		return false
	}

	ok := false
	for _, x := range sliceX {
		for _, y := range sliceY {
			// x found in y
			if x == y {
				ok = true
				break
			}
		}
		// x not found in y
		if !ok {
			return false
		}
	}

	return true
}
