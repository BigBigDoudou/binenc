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
	if !equalSlices(want, decoded) {
		t.Fatalf(`Decoded(int64(42)) = %v, want %v`, decoded, want)
	}
}

// Test if slices contains the same strings (not ordered)
func equalSlices(s1 []string, s2 []string) bool {
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
