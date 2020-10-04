package rewindable

import (
	"bufio"
	"strings"
	"testing"
)

func TestRewind(t *testing.T) {
	srcReader := strings.NewReader(`1
2
3
4
5`)
	r := NewReader(srcReader)
	sc := bufio.NewScanner(r)
	if !sc.Scan() || sc.Text() != "1" {
		t.Fatal("failed: 1st `1`")
	}
	r.Rewind()
	sc = bufio.NewScanner(r)
	if !sc.Scan() || sc.Text() != "1" {
		t.Fatal("failed: 2nd `1`")
	}
	if !sc.Scan() || sc.Text() != "2" {
		t.Fatal("failed: 2nd `2`")
	}
	r.Rewind()
	sc = bufio.NewScanner(r)
	if !sc.Scan() || sc.Text() != "1" {
		t.Fatal("failed: 3rd `1`")
	}
	if !sc.Scan() || sc.Text() != "2" {
		t.Fatal("failed: 3rd `2`")
	}
	if !sc.Scan() || sc.Text() != "3" {
		t.Fatal("failed: 3nd `3`")
	}
}
