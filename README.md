# haikunator [![Build Status](https://travis-ci.org/taion809/haikunator.svg?branch=master)](https://travis-ci.org/taion809/haikunator)

Golang port of haikunator

## Installation
`go get github.com/taion809/haikunator`

## Example
```go
package main

import (
	"fmt"
	"github.com/taion809/haikunator"
)

func main() {
	var haiku string

	h := haikunator.NewHaikunator()

	// Default haiku
	haiku = h.Haikunate()
	fmt.Println("Default: ", haiku) // => "long-smoke-5866"

	// Token range
	haiku = h.TokenHaikunate(100000)
	fmt.Println("With Token: ", haiku) // => "aged-pond-74896"

	// Don't include token
	haiku = h.DelimHaikunate("-")
	fmt.Println("With Delim: ", haiku) // => "wandering-water"

	// Use different delim without a token
	haiku = h.DelimHaikunate(".")
	fmt.Println("With Delim: ", haiku) // => "wandering.water"
	
	// Change token range and delimiter
	haiku = h.TokenDelimHaikunate(100000, ".")
	fmt.Println("With Token and Delim: ", haiku) // => "holy.pine.60124"

	// No token and space delimiter
	haiku = h.TokenDelimHaikunate(0, " ")
	fmt.Println("With Token and Delim: ", haiku) // => "holy pine"
}
```

## Documentation
Godoc can be found [here.](http://godoc.org/github.com/taion809/haikunator)

## Other Languages

Haikunator is also available in other languages. Check them out:

- Ruby: https://github.com/usmanbashir/haikunator
- Node: https://github.com/AtroxDev/haikunatorjs
- Python: https://github.com/AtroxDev/haikunator
- Go: https://github.com/yelinaung/go-haikunator

## License
This project is licensed under the MIT License.
