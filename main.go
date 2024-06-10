package main

import (
	"fmt"
	"profanityfilteritaapi/filters"
)

func main() {
	fmt.Println("hello")
	a := filters.PhraseFilter("Io non credo che uno stronzo come lui possa fare una figura di merda simile, per quanto coglione uno possa essere.")
	fmt.Println(a)
}
