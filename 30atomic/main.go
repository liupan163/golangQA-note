package _0atomic



try to add something on v0.1.1
continue test on v0.1.1
begin to add s.t. on v0.1.2
add tag v0.1.3
//原子操作

//Q:比较并交换操作与交换操作相比有什么不同？优势在哪？
//A:所谓比较并交换，是指，把新变值赋给变量，并返回变量的旧值。
//函数会先判断被操作的当前值，与我们预期的旧值是否相等。如果相等，并返回true表示交换已经进行。
//否则，就进行交换操作，返回false

//go在1.4版本中加入了一个sync/atomic包中添加了一个新类型Value。相当于一个容器

//atomic.Value是开箱即用的，两个指针方法---Store和Load。

//1、一旦atomic.Value值类型被真正的使用，就不能再被复制了。
//2、向原子值存储的第一个值，决定了他后面唯一能存的类型。

/*
var box6 atomic.Value
v6 := []int{1, 2, 3}
box6.Store(v6)
v6[1] = 4 // 注意，此处的操作不是并发安全的！
*/
//由于v6是切片类型的，所有的操作会影响到对他引用的地方.

/*
store := func(v []int) {
	replica := make([]int, len(v))
	copy(replica, v)
	box6.Store(replica)
}
store(v6)
v6[2] = 5 // 此处的操作是安全的。
*/
//切片复制，此时对v6的操作，不会影响到store结构体


//Q:如果要对原子值和互斥锁进行二选一，你认为最重要的三个决策条件应该是什么？
//A:  原子类可能有ABA问题。
//    若业务对ABA敏感，用锁。