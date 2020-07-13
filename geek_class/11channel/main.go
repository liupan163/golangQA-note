package _1channel

var uselessChan = make(chan<- int, 1)
//此处的chan只能接受，单向通道
//主要用处：限制代码
type Notifier interface {
	SendInt(ch chan<- int)
}

//函数声明的结果队列中使用单向通道
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

//函数getIntChan会返回一个<-chan int 类型的通道，这意味着该通道的程序，只能从通道中接受元素值。（实际就是一种对函数调用方的约束）

/*
var intChan2 = getIntChan()
for elem := range intChan2{
fmt.Printf("The element in intChan2: %v\n", elem)
}
*/

//专门为通道设置的select语句
/*
intChannels := [3]chan int{
	make(chan, 1),
	make(chan, 1),
	make(chan, 1)
}
index := rand.Intn(3)
intChannels[index] = index;
select {
	case <- intChannels[0]:
	case <- intChannels[1]:
	case _,ok<- intChannels[2]:
		if !ok{
			break;
		}
	default:
}
*/

//Q1、如果在select语句中发现某个通道已经关闭了，怎么屏蔽它所在分支？
//Q2、在select语句与for连用时，怎样直接退出外层的for语句？

