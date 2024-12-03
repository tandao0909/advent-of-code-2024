package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`\d+`)
	if (err != nil) {
		panic(err)
	}
	
}