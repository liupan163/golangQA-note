package interview_qa

import (
	"fmt"
	"runtime"
	"sync"
	"errors"
	"reflect"
)

func main() {
	//deferCall()
	//parseStudent()
	//closePackage()
	//imp2Interface()
	//deferFunc()
	//appendSliceFunc()
	//panicOrderDetailTest()
	//relectEg()
	//reflectDemo()
}
func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

//后中前  触发异常
//----------------------------------------------------------

type student struct {
	Name string
	Age  int
}

func parseStudent() {
	m := make(map[string]*student)
	students := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 24},
		{Name: "liu", Age: 24},
	}
	for _, stu := range students {
		m[stu.Name] = &stu
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

// foreach处都用的是副本。所以m[stu.Name]和&stu实际上指向同一个指针
//类闭包问题
//----------------------------------------------------------

func closePackage() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A:", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B:", i)
		}(i)
	}
	wg.Wait()
}

//闭包，值拷贝保存
//----------------------------------------------------------

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowA() {
	fmt.Println("showA")
}
func mainShowA() {
	t := Teacher{}
	t.ShowA()
}

//showA
//showA
//类自子类没有，往上去找父类方法
//----------------------------------------------------------

func chanRandomSelect() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chain := make(chan string, 1)
	int_chan <- 1
	string_chain <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println("value=", value)
	case value := <-string_chain:
		fmt.Println("value=", value)
	}
}

//select处理chan的时候，如果有多个符合条件,随便选择处理。即
//----------------------------------------------------------

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func deferParam() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

//10 1 2 3
//20 0 2 2
//2 0 2 2
//1 1 3 4
//defer顺序 && 参数函数顺序
//----------------------------------------------------------

func appendFunc() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

//输出[0 0 0 0 0 1 2 3]
//----------------------------------------------------------

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

//可能出现concurrent map read and map write
//----------------------------------------------------------
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- elem
			fmt.Println("Iter:", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func ChainMain() {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	v := <-th.Iter()
	fmt.Println("v:", "ch", v)
}

//问题处: ch:=make(chan interface{})
//原因:chan通道会缓存，寄会阻塞
//----------------------------------------------------------

type Human interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func imp2Interface() {
	//var human Human  = Stduent{}
	var human = Stduent{}
	think := "bitch"
	fmt.Println(human.Speak(think))
}

//方法集实现，只影响接口实现和方法表达式。
//----------------------------------------------------------

type People1 interface {
	Show()
}
type Student1 struct{}

func (stu *Student1) Show() {}

func live() People1 {
	var stu1 *Student1
	return stu1
}

func main2() {
	if live() == nil {
		fmt.Println("AAAAAA")
	} else {
		fmt.Println("BBBBBB")
	}
}

//接口类型判断 "BBBBBB"
//var stu1 *Student1  类型非空，值为空
//----------------------------------------------------------

func ISwitch() {
	//i := GetValue()
	//switch i.(type) {
	//case int:
	//	println("int")
	//case string:
	//	println("string")
	//case interface{}:
	//	println("interface")
	//default:
	//	println("unknown")
	//}
}
func GetValue() int {
	return 1
}

//问题点：i.(type)  √1、只能用在switch 2、i只能是interface类型
//----------------------------------------------------------

//----------------------------------------------------------//----------------------------------------------------------
func deferFunc() {
	//println(DeferFunc1(1))
	//println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

//deferFunc1的作用域是整个函数，返回4
//deferFunc2的作用域是函数，返回1
//deferFunc3的作用域是    ，返回3
//----------------------------------------------------------//----------------------------------------------------------

func appendSliceFunc() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

//切片的...
//----------------------------------------------------------

func nilCheck() {
	var x *int = nil
	Foo(x)
}
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
	} else {
		fmt.Println("non-empty interface")
	}
}

//接口类型为nil判断条件
//----------------------------------------------------------

func iotaEg() {
	const (
		x = iota
		y
		z = "zz"
		k
		p = iota
	)

}
func iotaCheck() {
	fmt.Println("iotaEg=>", iotaEg)
}

// 0 1 zz zz 4
//----------------------------------------------------------

func goFuncArea() {

	for i := 0; i < 10; i++ {
		//loop:
		println(i)
	}
	//goto loop
}

//goto函数作用域，goto函数不能进入函数内层代码

func typeAlias() {
	type MyInt1 int
	type MyInt2 = int
	var i int = 9
	//var i1 MyInt1 = i
	var i2 MyInt2 = i
	//fmt.Println(i1,i2)
	fmt.Println(i2)
}

//MyInt2为类型alias，能直接赋值
//MyInt1为definition，不能直接赋值
//----------------------------------------------------------

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func funcAreaTest() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}

//nil nil
//原因result, err := tryTheThing()  此处开始的err会遮罩函数作用域的err变量。两个err不相关
//----------------------------------------------------------

func panicOrderTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

func panicOrderDetailTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("++++")
			f := err.(func() string)
			fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	panic("panic")
}

//recover接受panic函数
//----------------------------------------------------------
func relectEg() {
	var num float64 = 1.2345

	typeInfo := reflect.TypeOf(num)
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	changeValue := pointer.Elem()
	changeValue.SetFloat(1.000)
	fmt.Println("type of pointer", changeValue.Type())
	fmt.Println("typeInfo", typeInfo, ",pointer", pointer, ",value", value)
	fmt.Println("value.Canset", value.CanSet())
	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

// 点：
// 只有指针类型的才可以.Elem()， .CanSet()和 .setFlocat()
// realValue := value.Interface().(已知的类型)
// typeInfo := relect.TypeOf(num)
// valueInfo := relect.ValueOf(num)
//
//----------------------------------------------------------

type ReflectStruct struct {
	Id   int
	Name string
	Age  int
}

func (u ReflectStruct) ReflectCallFunc() {
	fmt.Println("this is ReflectCallFunc")
}
func reflectDemo() {
	reflectDemo := ReflectStruct{1, "parker", 18}
	doReflectMethod(reflectDemo)
}

func doReflectMethod(demo ReflectStruct) {
	getType := reflect.TypeOf(demo)
	getValue := reflect.ValueOf(demo)
	fmt.Println("getType=", getType, ",getValue=", getValue)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v  = %v\n", field.Name, field.Type, value)
	}
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
