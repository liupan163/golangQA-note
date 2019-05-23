package _3struct

import "fmt"

//go语言中还有其他几样可以代表“指针”。其中最贴切的当属uintptr类型。该类型实际上是一个数值类型，也是go语言内置数据类型之一。

//golang语言标准库里面的unsafe包。unsafe包里面有一个类型叫做pointer，也代表指针。
//unsafe.Pointer可以表示任何指向可寻址的值的指针，同时它也是前面提到的指针值和uintptr值之间的桥梁---》通过它，这两种值可以相关转换。

func main() {
	//var t *int //t是个指针类型，初始化就是nil			fmt.Println("*t===", *t)报错
	t := new(int) //此时t指向 : ***"地址"***
	dd(t)

	var aPot *string
	/*
		理解&aPot  指针的地址
			*aPot  值
			aPot   指针变量
	*/
	fmt.Printf("&aPot: %p %#v\n", &aPot, aPot) // 输出 &aPot: 0xc42000c030    (*string)(nil)
	var aVar = "123"

	//指针声明，必须先赋值。
	//*aPot = "This is a Pointer" // 报错： panic: runtime error: invalid memory address or nil pointer dereference

	aPot = &aVar                                           //分配内存
	fmt.Printf("&aPot: %p %#v %#v \n", &aPot, aPot, *aPot) // 输出 &aPot: 0xc42000c030 (*string)(0xc42000e240) "123"
	*aPot = "This is a Pointer"
	fmt.Printf("&aPot: %p %#v %#v \n", &aPot, aPot, *aPot) // 输出 &aPot: 0xc42000c030 (*string)(0xc42000e240) "This is a Pointer"
}

func dd(t *int) {
	fmt.Println("t===", t)
	fmt.Println("*t===", *t)
}

//可寻址的addressable
//列举出go语言中哪些值是不可寻址的？
/*
	常量的值
	基本类型值的字面量
	算数操作的结果值
	对于各种字面量的索引表达式和切片表达式的结果值。 例外，切片字面量的索引结果值却是可寻址的。
	对字符串变量的索引表达式和切片表达式的结果值。
	对字典变量的索引表达式的结果值。
	函数字面量和方法字面量，以及对它们的调用表达式的结果值。
	•结构体字面量的字段值，也就是对结构体字面量的选择表达式的结果值。
	•类型转换表达式的结果值。
	•类型断言表达式的结果值。
	•接收表达式的结果值
*/

/*
	问题解析
	调用dog.SetName实际上是---> (&dog).SetName("monster")
	所以不能new("little pig").SetName("monster")
*/
