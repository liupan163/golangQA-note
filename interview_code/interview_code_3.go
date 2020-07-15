package main

import (
	"encoding/json"
	"fmt"
	"log"
	"unicode/utf8"
)

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
		fmt.Println(arr) //7 2 3
	}(x)
	fmt.Println(x) // 1 2 3
}
func arrayParameter2() {
	x := [3]int{1, 2, 3}
	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr) //7 2 3
	}(&x)
	fmt.Println(x) // 7 2 3
}

//切片作为参数
func sliceParameter3() {
	x := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr) // 7 2 3
	}(x)
	fmt.Println(x) // 7 2 3
}

//string类型是常量，不可更改。    切换成slice，再更改。
// 注意区分： 1、 []byte  2、[]rune
func stringConvert() {
	x := "text"
	xBytes := []byte(x)
	xBytes[0] = 's'

	xRunes := []rune(x)
	xRunes[0] = '谁'
}

//字符串长度
func stringLength() {
	testStr := "这是多类型hello world"
	fmt.Println(utf8.RuneCountInString(testStr)) //并不绝对正确，如单词：cliche
	testByte := []byte(testStr)
	fmt.Println(utf8.RuneCount(testByte))
}

//log.Fatal 和 log.Panic 不只是log, 会中断代码。
func stopTest() {
	log.Fatal("log fatal test, log.Fatal")
	log.Panic("log panic test, log.Panic")
}

//原子操作   goroutine和channel，或者sync包里面的锁

//range迭代 字符串
//如果字符串中，有任何非utf8格式的数据，应该先转换成byte slice，再进行操作
func stringRangeTest() {
	data := "A\xfe\x02\xff\x04"
	for _, v := range data {
		fmt.Printf("%#x ", v) // 错误用法 prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)
	}
	fmt.Println()
	for _, v := range []byte(data) {
		fmt.Printf("%#x ", v) //prints: 0x41 0xfe 0x2 0xff 0x4 (good)
	}
}

//range 迭代map，kep是乱序的

//switch的case情况，加fallthrough可强行执行，下个case语句

// 取反操作符 ^      同时也是亦或操作符

//运算符的优先级

//不可导出的struct字段(小写字母开头)，无法被encode. 	即：json、xml、gob等格式encode操作时，私有忽略
func encodeTest() {
	type MyData struct {
		One int
		two string
	}
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in) //prints main.MyData{One:1, two:"two"}

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) //prints {"One":1}

	var out MyData
	json.Unmarshal(encoded, &out)

	fmt.Printf("%#v\n", out) //prints main.MyData{One:1, two:""}
}
