# haikunator [![Build Status](https://travis-ci.org/taion809/haikunator.svg?branch=master)](https://travis-ci.org/taion809/haikunator)

Golang port of haikunator

## Example
```go
package main

import (
	"fmt"
	"github.com/taion809/haikunator"
)

func main() {
	h := haikunator.NewHaikunator()
	haiku := h.Haikunator()

	fmt.Println(haiku)
}
```

## License
Mit: See LICENSE for details
