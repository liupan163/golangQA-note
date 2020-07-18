package main

import "fmt"

/*
defer、return、返回值三者的执行逻辑应该是：
	return最先执行，return负责将结果写入返回值中；
	接着defer开始执行一些收尾工作；
	最后函数携带当前返回值退出。

defer  延迟的函数参数就已经计算完成了

执行顺序:先defer存参数，执行return语句，执行defer语句
1、注意返回值，定义位置
2、执行defer语句注意（参数是 值传递，还是引用传递）
*/

func main() {
	//fmt.Println("tt1 return=", tt1())
	//fmt.Println("tt2 return=", tt2())
	//fmt.Println("tt3 return=", tt3())
	//fmt.Println("tt4 return=", tt4()())
	fmt.Println("tt5 return =", tt5())
	//fmt.Println("tt6 return =", tt6())
}
/*
注意：如果defer后面只有一条语句，则其中的变量会立刻被赋值；如果defer后面是一个函数，则其中的变量会在执行时才被赋值。

*/
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
func tt3() int { //注意int 定义位置
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
		fmt.Println("defer tt4 after:", i)
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
