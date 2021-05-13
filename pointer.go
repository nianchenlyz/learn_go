package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指向 i，指向
	fmt.Println(*p) // 通过指针读取 i 的值,读取操作
	*p = 21         // 通过指针设置 i 的值，设置
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}
