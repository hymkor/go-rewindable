// +build ignore

package main

import (
	"bufio"
	"os"

	"github.com/zetamatta/go-rewindable"
)

func main() {
	reader := rewindable.NewReader(os.Stdin)
	for i := 0; i < 4; i++ {
		sc := bufio.NewScanner(reader)
		j := 0
		for sc.Scan() && j <= i+1 {
			println(">", sc.Text())
			j++
		}
		println("---")
		reader.Rewind()
	}
}
