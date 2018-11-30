package _4interface

//接口组合，go中的ReadWriteClose和ReadWriter接口就是组合的例子

//接口，一个指向动态值，一个指向类型信息。
//参照反射时候得用法--->reflect.Type和reflectValue



//Q：把一个值为nil的某了类型的变量，赋值给接口变量，那么在这个接口变量上仍然可以调用该接口的方法吗？ 需要注意什么？
//A：值方法不能用会报空指针
/*	空指针异常
	func (p *Person)TurnOnChrome(){
	}
	//引用类型--->ok
	func (p Person)TurnOnComputer(){}
	//属性可以调用，为空nil
*/

//ps:前提，golang里面全是浅拷贝