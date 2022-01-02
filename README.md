# Golang Word Reference Package
This is a third-party Golang package used to look up words and examples against WordRerefence

## HOW TO
```
package main

import (
	"fmt"

	"github.com/justinsowhat/wordreference-golang/client"
)

func main() {

	client := client.WordReferenceClient{
		Dict: "fren",
	}

	result := client.LookUpWord("de toute façon")

	fmt.Println(result)
}
```
