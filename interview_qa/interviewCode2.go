package main

import "fmt"

func init() {

}

/*
	defer 跟 return
*/
func main() {
	fmt.Println("tt1 return=", tt1())
	fmt.Println("tt2 return=", tt2())
	fmt.Println("tt3 return=", tt3())
	fmt.Println("tt4 return=", tt4()())
	fmt.Println("tt5 return =", tt5())
	fmt.Println("tt6 return =", tt6())
}

//return 1,defer tt1 0
func tt1() int {
	var i = 0
	defer fmt.Println("defer tt1", i)
	i++
	return i
}

//return 1,defer tt2 0
func tt2() int {
	var i = 0
	defer func(i int) {
		fmt.Println("defer tt2", i)
	}(i) //值传递
	i++
	return i
}

//return 1,defer tt3 1
func tt3() int {
	var i = 0
	defer func() {
		fmt.Println("defer tt3", i) //引用变量
		i++                         //2
	}()
	i++
	return i
}

//return :func(){ return i}, defer tt4: 1
//func()(): 2
func tt4() func() int {
	var i = 0
	defer func() {
		fmt.Println("defer tt4:", i)
		i++
	}()
	i++
	return func() int {
		return i //引用变量
	}
}

//return 13,defer 12
func tt5() (num int) {
	defer func() {
		fmt.Println("defer tt5:", num)
		num++ //引用变量
	}()
	return 12
}

//return 1 ,defer tt6: 0, defer tt6 after: 2
func tt6() (num int) {
	defer func(num int) {
		fmt.Println("defer tt6:", num)
		num++
		num++
		fmt.Println("defer tt6 after", num)
	}(num) //值传递
	num++
	return
}
