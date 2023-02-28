package main

import "fmt"

func main() {
	var (
		num  int
		num1 int
		num2 int
	)
	num = 100
	fmt.Printf("num的值为: %d, num的内存地址为: %p\n", num, &num)

	num = 200
	fmt.Printf("num的值为: %d, num的内存地址为: %p\n", num, &num)

	num1 = 123
	fmt.Printf("num1的值为: %d, num1的内存地址为: %p\n", num1, &num1)

	num2 = 456
	fmt.Printf("num2的值为: %d, num2的内存地址为: %p\n", num2, &num2)

}
