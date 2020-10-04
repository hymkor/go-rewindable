package rewindable

import (
	"bufio"
	"io"
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

func TestSeekStart(t *testing.T) {
	srcReader := strings.NewReader(`1
2
3
4
5`)
	var r io.ReadSeeker = NewReader(srcReader)
	r.Seek(8, io.SeekStart)
	sc := bufio.NewScanner(r)
	if !sc.Scan() || sc.Text() != "5" {
		t.Fatal("failed: SeekStart forward")
	}
	r.Seek(4, io.SeekStart)
	sc = bufio.NewScanner(r)
	if !sc.Scan() || sc.Text() != "3" {
		t.Fatal("failed: SeekStart backward")
	}
}
