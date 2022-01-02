# Golang Word Reference Package
This is a third-party Golang package used to look up words and examples against WordRerefence

## HOW TO
```
package main

import (
	"fmt"

	wd "github.com/justinsowhat/wordreference-golang"
)

func main() {

	client := wd.client.WordReferenceClient{
		Dict: "fren",
	}

	result := client.LookUpWord("de toute façon")

	fmt.Println(result)
}
```
