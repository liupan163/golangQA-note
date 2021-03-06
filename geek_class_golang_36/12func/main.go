package _2func

//高阶函数
//1接受其他函数作为参数传入
//2把其他函数作为结果返回

//Q:如何实现闭包？

//动态的生成一部分程序逻辑。根据生成功能不同的函数，继而影响后续的程序行为

/*
complexArray1 := [3][]string{
	[]string{"d", "e", "f"},
	[]string{"g", "h", "i"},
	[]string{"j", "k", "l"},
}
*/

//Q 1、complexArray1别传入函数的话，这个函数中对该参数值得修改会影响到它的原值吗？
//Q 2、函数真正拿到的参数值其实只是他们的副本，那么函数返回给调用方的结果值也会被复制吗？


//A1:golang中元素都是浅拷贝。若修改到了数组切片里面的元素，会影响源数组。
//A2:当函数返回指针类型的不会发生拷贝，非指针类型的会拷贝。
