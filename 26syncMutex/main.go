//golang自带的标准库中核心的代码包。

互斥量mutual exclusion  简称mutex
sync包中的Mutex就是对应类型。

time.Ticker类型。

互斥锁和读写锁的指针类型都是实现了 Lock这个接口

获取读写锁中的读锁？
变量.Rlock()