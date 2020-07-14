package main

import "fmt"

// http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html

func main() {

}

//nil 是interface、functions、pointers、map、slices、channels的默认零值。
func strongTypeEg() {
	//var x = nil //error     初始化要给出类型，golang是强类型语言
	var y interface{} = nil

	//_ = x
	_ = y
}

//允许对值为nil的slice添加值，不能给map添加值
func value2MapAndSlice() {
	var m map[string]int
	m["one"] = 1 //会报错，nil map
	// mm := make(map[string]int) //分配了内存空间，可以用来增删

	var s []int
	s = append(s, 1)
}

//array类型的值，作为函数参数
func arrayParameter() {
	x := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println(x)
}
func arrayParameter2() {
	x := [3]int{1, 2, 3}
	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr)
	}(&x)
	fmt.Println(x)
}
//切片作为参数
func sliceParameter3() {
	x := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println(x)
}

