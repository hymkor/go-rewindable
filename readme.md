go-rewindable
=============

go-rewindable is the package that add io.Reader .Rewind() method.
We can read io.Reader as many times as you like.

```go
package main

import (
    "bufio"
    "strings"

    "github.com/zetamatta/go-rewindable"
)

func main() {
    srcReader := strings.NewReader(`1
2
3
4
5`)

    reader := rewindable.NewReader(srcReader)
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
```

```
$ go run example1.go
> 1
> 2
---
> 1
> 2
> 3
---
> 1
> 2
> 3
> 4
---
> 1
> 2
> 3
> 4
> 5
---
```
