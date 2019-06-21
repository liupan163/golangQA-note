package _interface

import "fmt"

//接口型函数 ：适用于只有一个函数的接口
//面向接口编程

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	fmt.Println("method is begining ")
	EachFunc(persons, selfDesign)
}

type Handler interface {
	Do(k, v interface{})
}
type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}
func selfDesign(k, v interface{}) {
	fmt.Println("selfDesign=>", k, v)
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f)) //二参数特征：面向接口编程
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

//-----------------------------------------------------------------------------
//duck type
func main1() {

	/*
		eg1:
			value, b := interface.(Type)
			value 是 Type 的默认实例；b 是 bool 类型，表明断言是否成立。
	*/
	s := Student{Grade: 1, Major: "English"}
	v, b := interface{}(s).(Student)
	fmt.Println("v=", v, ",b=", b)

	/*
		eg2:
			i.(type)只能用在switch语句中
	*/
	v1, err := testParse(1)
	fmt.Println("v1=", v1, ",err=", err)
}

func testParse(val interface{}) (value bool, err error) {
	if val != nil {
		switch v := val.(type) {
		case bool:
			return v, nil
		case string:
			switch v {
			case "eg1":
				return true, nil
			case "eg2":
				return false, nil
			}
		}
	}
	return false, nil
}

type Car struct {
	Color     string
	SeatCount int
}
type Student struct {
	Grade int
	Major string
}

//-------------method1
type ReaderFunc1 func(p []byte) (n int, err error)

func (f ReaderFunc1) Read(p []byte) (n int, err error) {
	return f(p)
}

//-------------replace method2
type ReaderFunc2 struct{}

func (f ReaderFunc2) Read() {
	//	do	Something
}
