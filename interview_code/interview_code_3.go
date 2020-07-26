package main

import (
	"encoding/json"
	"fmt"
	"log"
	"unicode/utf8"
	"time"
)

// http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html

func main() {
	fmt.Println("333")
	closeChannelTest()
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

//无缓冲的channel发送数据，只要receiver准备好了，就立马能接受数据。（容量，阻塞问题）

//向已经关闭的channel发送数据会造成panic
func closeChannelTest(){
	ch := make(chan int )
	done := make(chan struct{})
	for i:=0;i<3;i++{
		go func(idx int){
			select {
			case ch <- (idx+1)*2 :
			fmt.Println(idx,"send result")
			case <-done:
				fmt.Println(idx,"exiting")
			}
		}(i)
	}
	fmt.Println("result :",<-ch)
	close(done)
	time.Sleep(3*time.Second)
	fmt.Println("thread is closing =")
}

//利用死锁特性，开关channel
func deadLockChannel(){
	inch := make(chan int)
    outch := make(chan int)

    go func() {
        var in <- chan int = inch
        var out chan <- int
        var val int
        for {
            select {
            case out <- val:
                out = nil
                in = inch
            case val = <- in:
                out = outch
                in = nil
            }
        }
	}()
	go func() {
        for r := range outch {
            fmt.Println("result:",r)
        }
    }()
    time.Sleep(0)
    inch <- 1
    inch <- 2
    time.Sleep(3 * time.Second)
}

//函数传参数类型，值传递还是引用传递
//特殊  map 或者 slice传参数

//http长链接   及时关闭 或清理 http请求头

//json中的数字 解码为interface类型

//值比较  struct、array、slice、map
//用==的前提是，两个结构体都是可比较类型
//go功能库提供的。reflect的DeepEqual()//这不总适用于slice

//panice恢复
//

//slice 拷贝后再用

//创建新类型，不会继承原有属性的情况
type myMutes sync.Mutex
func testNewType(){
	var mtx myMutes
	mtx.Lock()//报错
}
//正确用法
type myRightMutes {
	sync.Mutex
}

//defer是在调用他的函数结束时执行，而不是语句块结束后执行。
//尤其是for循环中，注意return


//只要值是可寻址的，就可以在值上直接调用指针方法。
//不是所有的值都是可寻址的，如map类型，通过interface引用的变量

//并发与并行: 
//     并发是  指同时管理很多事情，很多事物同时对同一数据进行操作。
//     并行是  同时做很多事情，让不同的代码片段同时在不同的处理器上执行.