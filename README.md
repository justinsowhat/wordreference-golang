[![Go](https://github.com/justinsowhat/wordreference-golang/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/justinsowhat/wordreference-golang/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/justinsowhat/wordreference-golang.svg)](https://pkg.go.dev/github.com/justinsowhat/wordreference-golang)

# Golang Word Reference Package
This is a simple Golang package used to look up words and examples against WordRerefence

## Installation
```
go get -u github.com/justinsowhat/wordreference-golang
```

## Example & Usage
```
package main

import (
	"fmt"

	wd "github.com/justinsowhat/wordreference-golang"
)

func main() {

	client := wd.WordReferenceClient{
		Dict: client.FRENCH_ENGLISH,
	}

	result := client.LookUpWord("subir")

	fmt.Printf("IPA: %s\n", result.IPA)
	fmt.Printf("Principal Translations: %s\n", result.TranslationGroups)

}
```

These are the list of supported dictionaries on WordReference. Please note, not all of them are tested with this package.

```
    ENGLISH_SPANISH   
	SPANISH_ENGLISH        
	SPANISH_FRENCH         
	SPANISH_ITALIAN   
	SPANISH_GERMAN       
	SPANISH_PORTUGESE      
	ENGLISH_FRENCH           
	FRENCH_ENGLISH        
	FRENCH_SPANISH          
	ITALIAN_ENGLISH        
	ENGLISH_ITALIAN        
	ITALIAN_SPANISH      
	ENGLISH_GERMAN
	GERMAN_ENGLISH       
	GERMAN_SPANISH         
	ENGLISH_DUTCH        
	DUTCH_ENGLISH       
	ENGLISH_SWEDISH     
	SWEDISH_ENGLISH     
	ENGLISH_RUSSIAN        
	RUSSIAN_ENGLISH     
	ENGLISH_PORTUGESE       
	PORTUGESE_ENGLISH      
	PORTUGESE_SPANISH       
	ENGLISH_POLISH          
	POLISH_ENGLISH          
	ENGLISH_ROMANIAN       
	ROMANIAN_ENGLISH        
	ENGLISH_CZECH            
	CZECH_ENGLISH           
	ENGLISH_GREEK            
	GREEK_ENGLISH          
	ENGLISH_TURKISH        
	TURKISH_ENGLISH          
	ENGLISH_CHINESE       
	CHINESE_ENGLISH        
	ENGLISH_JAPANESE     
	JAPANESE_ENGLISH      
	ENGLISH_KOREAN        
	KOREAN_ENGLISH      
	ENGLISH_ARABIC    
	ARABIC_ENGLISH   
	ENGLISH_ICELANDIC  
```


## License
The MIT license is here[https://github.com/justinsowhat/wordreference-golang/blob/main/LICENSE].
