package search

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	in, _ := os.Open("testdata/missing_number.in.txt")
	bin := new(bytes.Buffer)
	io.Copy(bin, in)

	out := new(bytes.Buffer)

	MissingNumber(bin, out)

	exp, _ := os.Open("testdata/missing_number.exp.txt")
	expBuf := new(bytes.Buffer)
	io.Copy(expBuf, exp)

	if strings.TrimSpace(expBuf.String()) != strings.TrimSpace(out.String()) {
		t.Errorf("exp: %s, got: %s", expBuf.String(), out.String())
	}
}
