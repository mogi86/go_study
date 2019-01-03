package variable_length

import "fmt"

//可変長引数
func PrintVariableLength(who ...string) {
	fmt.Println(who)
}

func PrintRange(who ...int) {
	for i, v := range who {
		fmt.Print(fmt.Sprintf("index:%d,value:%d\n", i, v))
	}
}
