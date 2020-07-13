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


