（判断一个操作是否是原子的，可以使用go run race 命令做数据的竞争检测）


1、GOROOT、GOPATH和GOBIN

GOROOT：就是golang的安装路径

Q:设置GOPATH有什么意义？

A:环境变量GOPATH的值可以是一个目录的路径，也可以包含多个目录路径，每个目录都代表Go语言的一个工作区。
这些工作区用于放置Go语言的源码文件，以及安装后的归档文件（archive file，就是以.a为扩展名的文件）和可执行文件。

pkg子目录---归档文件；
bin子目录---可执行文件；
src---源码路径。

Q:go mod 和 go vendor用法？

A:

2、flag代码包
go语言标准库中---专门用于接受和解析命令参数。

多用这个方式
flag.StringVar(&name,"name","everyone","The greeting object")
flag.StringVar(&GCmd.ConfigPath, CmdConfigPath.String(), "", "config path")
flag.BoolVar(&help, "h", false, "")
参数说明：
- 第1个用于存储该命令参数值得地址。变量name的地址，&name
- 第2个参数为该命令指定命令参数名称，name。
- 第3个指定默认值
- 第4只作为简短说明，打印命令说明时会用到。

区别于flag.String。  后者会返回个分配好存储命令参数地址。
用法示例：
var name string 改动后。
var name = flag.String("name","everyone","The greeting object")

3、实体访问权限
大小写

4、实体
golang中，包括变量、常量、函数、接口、结构体。
golang是静态语言，一旦类型确定，不可改。
变量类型：自定义、程序定义类型

5、类型断言
语法是：    x.(T)
示例代码：
interface{}代表空接口，任何类型都是她的实现
value, ok := interface{}(container).([]string)

6、类型转换
语法是：    T(x)
用法示例:
uint8(255)

- 整数转字符串是可行的，但是得在范围内，即一个有效的unicode码
- a）string在转成utf-8编码的字符串时候，会被拆分成零散、独立的字节
除了与ASCII兼容的部分，剩余部分会被拆分成单一字节
- b）string在转成[]rune时候，代表着字符串会被拆分成一个个Unicode字符

7、别名类型？潜在类型
type MyString = string //Type Aliases
type MyString2 string  //无等号

8、切片


9、容器container、链表container/list、循环链表(环)container/ring

container/list包含两个实体：List和Element，前者是双向链表
用法示例：
    l:=list.New()
    e2 := l.PushFront(2)
    e3 := l.PushBack(3)
    l.InsertBefore(3,e2)
    l.InsertAfter(2,e3)

container/ring包含两个实体：next和Value
用法示例：
    r:= ring.New(10)
    r.value=100
    r = r.next()

Ring与List的区别在哪儿？
/*
container/ring包中的Ring类型实现的是一个循环链表---环。  List内部就是一个循环链表。他的根元素永远不会持有任何实际的元素值，
该元素的存在，就是为了链接这个循环表的收尾两端。
所以说，List的零值是一个包含了根元素，但不包含任何元素的空链表。

1、Ring类型的数据结构仅由他自身即可代表，而List类型则需要它以及Element联合表示。
2、Ring类型的值严格来讲，代表了其所属循环链表中的一个元素，而List代表一个完整的链表。
3、创建Ring值得时候，我们可以指定他包含元素的数量。List不能这么做。
4、var r ring.Ring 长度为1，而List长度为0。List根元素不会持有实际元素值。
5、Ring值得len算法复杂度是O（N），而List值得Len方法的算法复杂度是O（1）。
*/

使用场景：list用在FIFO场景
          ring使用在定长队列，如轮播等。
          list可以作为queue和stack的基础数据结构
          ring可以用来保存固定数量的元素，如：最近100条日志

10、字典类型Map
map是一个引用类型。 就是个 哈希表hash table
key类型受限，（键不能是，函数、切片、map）
value可以是任意类型

哈希碰撞问题：每个键都转换成哈希值，golang会对比hash值和key值

并发：不安全。
（判断一个操作是否是原子的，可以使用go run race 命令做数据的竞争检测）
解决方案： sync.Map.  （相当于表级别锁）
    通过读写分离实现，降低锁时间来提高效率。
    缺点：不适合大量写的场景。
适合场景：大量读，少量写。
优化方向参考：https://github.com/orcaman/concurrent-map/blob/master/README-zh.md

11、通道channel（核心）
- 进入通道的值，是元素的副本。
- 通道一旦关闭，再对他进行操作，就会引发panic
- 关闭一个已经关闭了的通道，也会引发panic
- 不要从接收端关闭channel
用法示例：
    res,boolean <-ch
    第二个boolean表达式，来判断是否获取成功

Q:通道的长度代表着什么？什么时候通道的长度和容量相同？
A:阻塞的时候。容量--->代表channel缓存的长度。

Q:元素在通达被复制时，深复制还是浅复制？
A:golang里面没有深复制。数组是值类型，所以会被完全复制。

select--配合channel使用
用法示例：
select {
	case <- intChannels[0]:
	case <- intChannels[1]:
	case _,ok<- intChannels[2]:
		if !ok{
			break;
		}
	default:
}

Q:在select语句与for连用时，怎样直接退出外层的for语句？
A:增加label代码块，退出for语句

Q:如果在select语句中发现某个通道已经关闭了，怎么屏蔽它所在分支？
A:候选分支中将接收表达式赋值给两个变量，使用第二个变量来判断 channel是否关闭，如果关闭，可以使用 break 来屏蔽这条候选分支的逻辑代码

12、高阶函数
把其他函数当做参数传进来，把函数当结果传出去
Q:参数传入函数的话，这个函数中对该参数值得修改，会影响到它的原值吗？
A:golang中元素都是浅拷贝。 若修改到了数组切片里面的元素，会影响源数组。 或者地址。

Q:函数真正拿到的参数值其实只是他们的副本，那么函数返回给调用方的结果值也会被复制吗？
A:当函数返回指针类型的不会发生拷贝，非指针类型的会拷贝。

13、结构体struct

嵌入字段（也叫匿名字段）。
定义：声明一个结构体的时候，只有字段名，没有字段名称。也就是说，只有名称，没有类型。
用法示例：
type Student struct{
    Human      //匿名字段
    special   String
}
mark:=Student(Human{"nameDDD",33,false},"ddd")
mark.name //mark可以直接引用Human里面的属性字段。

类似继承，利用字段嵌入，实现了  类型组合。

指针方法
func (cat *Cat) SetName(name string)  {
	cat.name = name
}
值方法
func (cat Cat) SetName(name string)  {
	cat.name = name
}

Q:我们可以在结构体类型中嵌入某个类型的指针类型吗？需要注意哪些？
A:可以在结构体内嵌入某个类型的指针类型，他和普通指针类似，默认初始化为nil。用之前需要初始化

Q:字面量struct{}代表了什么？有什么用处？
A:空结构体不占用空间，但是具有结构体的一切属性。如可以有方法，写入channel。


14、interface
接口分为type和value，一个指向动态值，一个指向类型信息。
参照反射类型就是，reflect.Type和reflect.Value

//Q：把一个值为nil的某了类型的变量，赋值给接口变量. 那么在这个接口变量上仍然可以调用该接口的方法吗？ 需要注意什么？
//A：值方法不能用,会报空指针
/*	空指针异常
	func (p *Person)TurnOnChrome(){
	}
	//引用类型--->ok
	func (p Person)TurnOnComputer(){}
	//属性可以调用，为空nil
*/

15、Pointer
unsafe包里面有一个类型叫做Pointer，通常用来代表指针。  unsafe是类型安全的操作
- 任何类型的指针都可以被转化为Pointer
- Pointer可以被转化为任何类型的指针
- uintptr可以被转化为Pointer
- Pointer可以被转化为uintptr

uintptr类型。该类型实际上是一个数值类型，也是go语言内置数据类型之一，也能代表指针

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

/*
	问题解析
	调用dog.SetName实际上是---> (&dog).SetName("monster")
	所以不能new("little pig").SetName("monster")
*/


16、atomic
原子操作:在sync/atomic包中，声明了很多相关函数
go的原子操作基于CPU和操作系统的。

trigger := func(i uint32, fn func()){
    for{
        if n := atomic.LoadUint32(&count); n==i {
            fn()
            atomic.AddUint32(&count,1)
            break
        }
        time.Sleep(time.Nanosecond)
    }
}
这儿trigger实现了自旋spinning，除非发现条件满足，否则会不断检查。

原子类型无法针对int类型的值来做，int宽度会根据计算机架构而改变int32或者int64


Q: sync/atomic包中提供了几种原子操作？可操作的数据类型又有那些？
A: add   CAS(compare and swap) load store swap

atomic.Value是开箱即用的，两个指针方法---Store和Load。
1、一旦atomic.Value值类型被真正的使用，就不能再被复制了。
2、向原子值存储的第一个值，决定了他后面唯一能存的类型。


//ps:针对unsafe.Pointer类型，该包并未提供原子加法操作的函数。

Q:比较并交换操作与交换操作相比有什么不同？优势在哪？
A:所谓比较并交换，是指，1把新变值赋给变量，并返回变量的旧值。
    函数会先判断被操作的当前值，与我们预期的旧值是否相等。如果相等=>并返回true表示交换已经进行。
                                                          否则，就进行交换操作，返回false
    atomic.CompareAndSwapUint32(&sum, 100, sum+1)

Q:如果要对原子值和互斥锁进行二选一，你认为最重要的三个决策条件应该是什么？
A:  原子类可能有ABA问题。
    若业务对ABA敏感，用锁。


17、mutex互斥量
mutual exclusion  简称mutex。      sync/Mutex类型

互斥锁和读写锁的指针类型都是实现了 Lock这个接口

time.Ticker类型。  定时器作用

18、条件变量sync.Cond

基于互斥锁的，必须有互斥锁才能发挥作用。
等待通知wait、单发通知signal、广播通知broadcast

示例用法:
var cond sync.Cond // 创建全局条件变量

// 生产者
func producer(out chan<- int, idx int) {
    for {
        cond.L.Lock()       // 条件变量对应互斥锁加锁
        for len(out) == 3 { // 产品区满 等待消费者消费
            cond.Wait() // 挂起当前协程， 等待条件变量满足，被消费者唤醒
        }
        num := rand.Intn(1000) // 产生一个随机数
        out <- num             // 写入到 channel 中 （生产）
        fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", idx, num, len(out))
        cond.L.Unlock()         // 生产结束，解锁互斥锁
        cond.Signal()           // 唤醒 阻塞的 消费者
        time.Sleep(time.Second) // 生产完休息一会，给其他协程执行机会
    }
}

//消费者
func consumer(in <-chan int, idx int) {
    for {
        cond.L.Lock()      // 条件变量对应互斥锁加锁（与生产者是同一个）
        for len(in) == 0 { // 产品区为空 等待生产者生产
            cond.Wait() // 挂起当前协程， 等待条件变量满足，被生产者唤醒
        }
        num := <-in // 将 channel 中的数据读走 （消费）
        fmt.Printf("---- %dth 消费者, 消费数据 %3d,公共区剩余%d个数据\n", idx, num, len(in))
        cond.L.Unlock()                    // 消费结束，解锁互斥锁
        cond.Signal()                      // 唤醒 阻塞的 生产者
        time.Sleep(time.Millisecond * 500) //消费完 休息一会，给其他协程执行机会
    }
}
func main() {
    rand.Seed(time.Now().UnixNano()) // 设置随机数种子
    quit := make(chan bool)          // 创建用于结束通信的 channel

    product := make(chan int, 3) // 产品区（公共区）使用channel 模拟
    cond.L = new(sync.Mutex)     // 创建互斥锁和条件变量

    for i := 0; i < 5; i++ { // 5个消费者
        go producer(product, i+1)
    }
    for i := 0; i < 3; i++ { // 3个生产者
        go consumer(product, i+1)
    }
    <-quit // 主协程阻塞 不结束
}

Q:*sync.Cond类型的值可以被传递吗？那 sync.Cond类型的值呢？
A:

19、new和make区别

//new和make  区别用法
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
//var p *[]int = new([]int)
//*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)

//new对象 跟 声明的区别
type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}
p := new(SyncedBuffer) // type *SyncedBuffer
var v SyncedBuffer     // type  SyncedBuffer

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

20、return 和 defer 顺序

21、iota用法
