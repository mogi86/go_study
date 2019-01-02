package construction

import "fmt"

type Vertex1 struct {
    X, Y int
}

type Vertex2 struct {
    X, Y int
}

func PrintStruct() {
    v := Vertex1{100, 1000}

    fmt.Println(v.Y)
}

func UsePointer() {
    v := Vertex2{2, 3}
    p := &v
    p.X = 777
    fmt.Printf("構造体: %v\n", v)
    fmt.Printf("構造体ポインタ: %v\n", p)
}
