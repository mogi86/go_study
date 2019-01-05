package maps

import "fmt"

func PrintMaps() {
    var m map[string]int
    m = make(map[string]int)

    m["hoge"] = 100
    m["fuga"] = 200

    fmt.Println(m)
}

func IsExists(key string) {
    var m map[string]int
    m = make(map[string]int)

    m["hoge"] = 100
    m["fuga"] = 200

    v, ok := m[key]
    fmt.Println(v)
    fmt.Println(ok)
}
