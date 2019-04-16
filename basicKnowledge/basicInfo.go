package basicKnowledge

import (
	"fmt"
	"os"
	"sync"
	"bytes"
)

func main() {
	/*
		基本类型
	*/
	var age int = 10
	time := 10
	const gender int = 10
	fmt.Println("age=", age, ",time=", time, ",gender", gender)
	/*
		类型转换
		golang对类型有严格限制
	*/
	var a int32 = 13
	var b int64 = 20
	c := int64(a) + b
	fmt.Println("a=", a, ",b=", b, ",c", c)
	/*
		map
	*/
	ages := make(map[string]int)
	ages["parker"] = 18
	for key, value := range ages {
		fmt.Println("key=", key, "value=", value)
	}
	/*
		函数   ---> fmt.print()
		&&方法 ---> func(p Person) GetName() stirng{}
	*/
	/*
		指针
	*/
	var age1 int = 10
	var p *int = &age1
	*p = 100
	fmt.Println("age=", age1)
	/*
		并发
	*/
	ipcBroadcast := make(chan int)
	go func() {
		sum := 0
		for i := 0; i < 10; i++ {
			sum = sum + i
		}
		ipcBroadcast <- sum
	}()
	fmt.Println(<-ipcBroadcast) //留坑，理解协程

	/*
		defer代替finally
	*/
	f, _ := os.Open("fileName")
	defer f.Close()

}

/*
	结构体替代类
*/
type Address struct {
	city string
}
type Person struct {
	Address
	age  int
	name string
}

//注意：方法名大小写
func (p Person) GetName() string {
	return p.name
}

/*
	接口
*/
type Greet interface {
	SayHello() string
}

func (p Person) SayHello() string {
	return p.GetName() + " say hello to u"
}

//new对象 跟 声明的区别
type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}
//p := new(SyncedBuffer) // type *SyncedBuffer
//var v SyncedBuffer     // type  SyncedBuffer

/*
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	f.dirinfo = nil
	f.nepipe = 0
	return f
}
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := File{fd, name, nil, 0}
	return &f
}
*/

//new和make  区别用法
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
//var p *[]int = new([]int)
//*p = make([]int, 100, 100)

// Idiomatic:
//v := make([]int, 100)
