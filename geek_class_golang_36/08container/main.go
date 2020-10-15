package _8container

//slice切片代表数组
//对应的链表---list

//golang语言中的链表在contianer中：List（双向链表）和Element（链表元素的结构）

//container中: ring循环链表

/*
list中有MoveBefore和MoveAfter方法，用于将给定的元素移动到另一个元素的前面和后面。
另外还有MoveToFront和MoveToBack方法。
func (l *List) MoveBefore(e, mark *Element)
func (l *List) MoveAfter(e, mark *Element)

func (l *List) MoveToFront(e *Element)
func (l *List) MoveToBack(e *Element)
*/

/*
List中，自己生成的Elem不会被链表接受。
List包含的方法中，用于插入新元素的方法，只接受interface{}类型的值。这些方法在内部会使用Element值包装接收到的新元素。

这么做主要原因是避免了链表内部关联遭到外界破坏
*/

/*
func (l *List) Front() *Element
func (l *List) Back() *Element

func (l *List) InsertBefore(v interface{}, mark *Element) *Element
func (l *List) InsertAfter(v interface{}, mark *Element) *Element

func (l *List) PushFront(v interface{}) *Element
func (l *List) PushBack(v interface{}) *Element
*/

//这些方法都会把一个Element值得指针作为结果返回，就是链表留给我们的安全的“接口”。 用这些指针再去做操作。

//一、为什么链表可以做到开箱即用？
/*
通过var l list.List声明的链表l可以直接使用的原因?
List的结构体类型有两个字段，一个是Element类型的root（跟元素），另个int类型的len（长度）.他们都是包级私有的

字段root和len都被赋予了相应的零值。len为0，root就是该类型的空壳，就是Element{}

Element类型：
包含了几个私有字段，分别用于存储前一个元素、后一个元素和所属链表的指针值。
另外还有一个value的公开的字段--->就是该元素的实际值, interface{}类型的。


......
list利用了自身，以及Element在结构上的特点，使链表可以开箱即用，也均衡了性能。*/

//二、Ring与List的区别在哪儿？
/*container/ring包中的Ring类型实现的是一个循环链表---环。  List内部就是一个循环链表。他的根元素永远不会持有任何实际的元素值，
该元素的存在，就是为了链接这个循环表的收尾两端。
所以说，List的零值是一个包含了根元素，但不包含任何元素的空链表。

1、Ring类型的数据结构仅由他自身即可代表，而List类型则需要它以及Element联合表示。
2、Ring类型的值严格来讲，代表了其所属循环链表中的一个元素，而List代表一个完整的链表。
3、创建Ring值得时候，我们可以指定他包含元素的数量。List不能这么做。
4、var r ring.Ring 长度为1，而List长度为0。List根元素不会持有实际元素值。
5、Ring值得len算法复杂度是O（N），而List值得Len方法的算法复杂度是O（1）。
*/

//list可以作为queue和stack的基础数据结构
//ring可以用来保存固定数量的元素，如：最近100条日志
//heap可以用来排序。  游戏编程中是一种高效的定时器实现方案。？？？？？？

