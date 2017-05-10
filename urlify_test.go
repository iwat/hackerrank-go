package algorithms

import (
	"testing"
)

func TestURLify(t *testing.T) {
	buffer := []byte("Mr John Smith    ")
	URLify(buffer, 13)
	if string(buffer) != "Mr%20John%20Smith" {
		t.Fatal(string(buffer))
	}

	buffer = []byte("Mr John Smith         ")
	URLify(buffer, 13)
	if string(buffer) != "Mr%20John%20Smith     " {
		t.Fatal("[", string(buffer), "]")
	}
}
