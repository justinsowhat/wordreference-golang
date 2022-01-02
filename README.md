# Golang Word Reference Package
This is a simple Golang package used to look up words and examples against WordRerefence

## Installation
```
go get -u github.com/justinsowhat/wordreference-golang
```

## Example
```
package main

import (
	"fmt"

	wd "github.com/justinsowhat/wordreference-golang"
)

func main() {

	client := wd.WordReferenceClient{
		Dict: "fren",
	}

	result := client.LookUpWord("subir")

	fmt.Printf("IPA: %s\n", result.IPA)
	fmt.Printf("Principal Translations: %s\n", result.PrincipalTranslations)
	fmt.Printf("Additional Translations: %s\n", result.AdditionalTranslations)
	fmt.Printf("Compound Forms: %s\n", result.CompoundForms)

}
```
