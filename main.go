package main

import (
	"fmt"

	"github.com/cheynewallace/tabby"
)

func main() {
	fmt.Println("hello world test")
	t := tabby.New()

	t.AddHeader("Name", "Title", "DEPARTMENT")
	t.AddLine("Kerry", "Developer", "Enginner")
	t.AddLine("Test", "Developer", "Whatever")
	t.AddLine("Sandra", "Developer", "testtest")
	t.Print()

}
