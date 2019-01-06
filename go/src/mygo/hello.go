package main

import (
	"fmt"
	"github.com/fatih/color"
	"mygo/array"
	"mygo/construction"
	"mygo/variable_length"
	"mygo/maps"
	"mygo/closure"
)

func main() {
	fmt.Println("Hello, 世界")
	color.Red("赤")

	//配列
	array.PrintArray()

	//構造体
	construction.PrintStruct()
	construction.UsePointer()

	//可変長引数
	variable_length.PrintVariableLength("hoge", "fuga")
	variable_length.PrintRange(1, 3, 5, 7, 9)

	//maps
	maps.PrintMaps()
	maps.IsExists("hoge2")

	//クロージャー
	func1 := closure.PrintClosure(100)
	func1()
	func1()
	func1()
	func2 := closure.PrintClosure(200)
	func2()
	func2()
	func2()
}
