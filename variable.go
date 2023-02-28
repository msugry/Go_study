package main

import "fmt"

//func main() {
//	var name string = "msugry"
//	var age int = 18
//	fmt.Println(name, age)
//}

func main() {
	/*批量定义变量*/

	//var (
	//	name string
	//	age  int
	//	addr string
	//)

	/*
		// var形式的声明语句往往是用于需要显式指定白你两类型地方，或者因为变量稍后会被重新赋值二初始值无关紧要的地方。
		// 当一个变量被声明后，如果没有显示的给他赋值，系统自动赋予它该类型的零值：
		// string 默认值为空
		// int 默认值为0
		// bool 默认值为false
		// 切片、函数、指针变量的默认为nil
	*/

	//name = "msugry"
	//age = 18
	//addr = "Beijing"

	/*
		变量初始化
	*/
	// 自动推导——短变量声明并初始化
	// go语言的推导声明写法，编译器会自动根据右值类型推断出左值的对应类型
	/*
		1、定义变量，同时显式初始化
		2、不能提供数据类型
		3、只能用在函数内部，不能随便到处定义
	*/
	name := "Msugry" //no new variables on left side of :=
	age := 18
	addr := "hello"
	fmt.Println(name, age, addr)
	fmt.Printf("name: %T, age: %T", name, age)
}
