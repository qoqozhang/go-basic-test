package type_alias_decl

import "fmt"

// 参数类型声明
type str string
type T[P any, S interface{ ~[]byte | ~string }] struct { // ~用来表示最底层的类型，如果不加~的花，例子t1会报错
	Description P
	Name        S
}

func (t T[P, S]) print() {
	fmt.Printf("%v %v\nl", t.Description, t.Name)
}

var t1 = T[string, str]{
	Description: "Jacky",
	Name:        "Mike",
}
var t2 = T[int, []byte]{
	Description: 123,
	Name:        []byte("Hello"),
}

// comparable 可比较的类型都可以用这个类型来表示
type TT[T comparable] struct {
	Compare T
}

var tt1 = TT[string]{Compare: "Hello"}
var tt2 = TT[int]{Compare: 123}
var tt3 = TT[bool]{Compare: true}
var tt4 = TT[byte]{Compare: byte(1)}
var tt5 = TT[rune]{Compare: rune(2)}
var tt6 = TT[float32]{Compare: float32(3)}
var tt7 = TT[struct{}]{Compare: struct{}{}}

func main() {
	t := new(T[string, string])
	t.print()
}
