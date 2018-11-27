package _6entity_03

//类型检查???

//一、类型断言--->表达式的语法形式是x.(T)
//interface{}代表空接口，任何类型都是她的实现
value, ok := interface{}(container).([]string)

//它包括了用来把container变量的值,转换成空接口值的interface{}(container)
//以及判断前者类型的是否为切片类型[]string的.([]string)

如果成功，赋值为value，ok为true。
否则ok为false, value--->null

//二、类型转换
//语法形式--->T(x)
//eg:uint8(255)
正数的补码等于原码，负数的补码才是反码＋1
整数转字符串是可行的，但是得在范围内，即一个有效的unicode码

a）string在转成utf-8编码的字符串时候，会被拆分成零散、独立的字节
除了与ASCII兼容的部分，剩余部分会被拆分成单一字节
b）string在转成[]rune时候，代表着字符串会被拆分成一个个Unicode字符

//三、别名类型？潜在类型？
type MyString = string //Type Aliases
type MyString2 string  //注意，这里没有等号。
