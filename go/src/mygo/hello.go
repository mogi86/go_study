package main

import "fmt"
import "github.com/fatih/color"
import "mygo/array"
import "mygo/construction"

func main() {
	fmt.Println("Hello, 世界")
	color.Red("赤")

	array.PrintArray()

	construction.PrintStruct()
	construction.UsePointer()
}
