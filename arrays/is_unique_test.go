package arrays

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	if IsUnique("the quick brown fox jumps over the lazy dog") {
		t.Fatal("Spaces is not unique")
	}
	if !IsUnique("thequickbrownfxjmpsvlazydg") {
		t.Fatal("It must be unique")
	}
}
