package main

import (
	"fmt"
	"os"
)

func main() {
	var name, age, name1, age1 string
	fmt.Scan(&name, &age) // 从stdin获取赋值给变量， 多个变量用空格，换行分割

	fmt.Sscan("zhang 11", &name1, &age1)

	fmt.Printf("name: %s, age: %s\n", name, age)
	fmt.Fprintf(os.Stdout, "name: %s, age: %s\n", name, age) // 通过Fprintf 把字符串写入到具有 io.Writer 的对象上

	w, _ := os.Create("./write.txt")
	defer w.Close()
	fmt.Fprint(w, "name: ", name1, ", age: ", age1) // 写入到具有 io.Writer 的w 上
}
