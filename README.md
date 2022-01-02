[![Go](https://github.com/justinsowhat/wordreference-golang/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/justinsowhat/wordreference-golang/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/justinsowhat/wordreference-golang.svg)](https://pkg.go.dev/github.com/justinsowhat/wordreference-golang)

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

## License
The MIT license is here[https://github.com/justinsowhat/wordreference-golang/blob/main/LICENSE].
