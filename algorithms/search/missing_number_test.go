package search

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	in, _ := os.Open("testdata/missing_number.in.txt")
	out := new(bytes.Buffer)

	bin := bufio.NewReader(in)
	MissingNumber(bin, out)

	var val int

	fmt.Fscanf(out, "%d", &val)
	if val != 7251 {
		t.Error("#0 != 7251")
	}
	fmt.Fscanf(out, "%d", &val)
	if val != 7259 {
		t.Error("#0 != 7259")
	}
	fmt.Fscanf(out, "%d", &val)
	if val != 7276 {
		t.Error("#0 != 7276")
	}
	fmt.Fscanf(out, "%d", &val)
	if val != 7279 {
		t.Error("#0 != 7279")
	}
	fmt.Fscanf(out, "%d", &val)
	if val != 7292 {
		t.Error("#0 != 7292")
	}
	fmt.Fscanf(out, "%d", &val)
	if val != 7293 {
		t.Error("#0 != 7293")
	}
}
