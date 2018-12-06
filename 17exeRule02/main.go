package _3struct
// 原子操作函数
/*
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
*/

//在sync/atomic包中声明了很多原子操作的函数
//这儿trigger实现了自旋spinning，除非发现条件满足，否则会不断检查。

//原子类型无法针对int类型的值来做，int宽度会根据计算机架构而改变int32或者int64
