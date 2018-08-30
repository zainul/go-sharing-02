package sampleMap

import (
	"fmt"
)

// TitleCode ...
var TitleCode map[string]int

func init() {
	titleCode := make(map[string]int)

	titleCode["Mr"] = 1
	titleCode["Mrs"] = 2

	TitleCode = titleCode
}

func main() {
	fmt.Println(choose("Mr"))
	fmt.Println(chooseMap("Mrs"))
}

func choose(title string) int {
	if title == "Mr" {
		return 1
	} else if title == "Mrs" {
		return 2
	}
	return 0
}

func chooseMap(title string) int {
	if val, found := TitleCode[title]; found {
		return val
	}
	return 0
}
