package main

import (
	"fmt"
	"root/color"
)

func main() {
	red := color.New(color.Bold, color.Underline)

	fmt.Println(red("Top"))
}
