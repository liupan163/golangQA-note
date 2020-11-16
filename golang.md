# go

## 反射

## 指针

## go runtime

## 缓存、限流、（降级）

# go并发处理
## 同步原语的适用场景
-   共享资源
-   任务编排
-   消息传递

## 并发原语Mutex
mutex是并发编程的基石，主要 1、接口Lock和UnLock接口，2、data race的临界区。
-   Mutex
    -   RWMutex
    -   锁竞争的实现，1、多给新人机会（在cpu中的线程） 2、避免竞争饥饿现象出现
    
可通过 [race detector](https://blog.golang.org/race-detector) 预先排除问题。具体运行 Go 代码的时候，加上race 参数。
也可通过vet工具，增加检测测试。

****tip：需要注意的点，一个goroutine中的lock方法，可以通过另一个goroutine来调用unlock方法****

## 锁功能扩展
### 可重入锁实现
Mutex本身不是可重入锁。手动实现一个。
核心：获取记录goroutine id。不同版本有不同方法，[获取goroutineId 工具库](https://time.geekbang.org/dashboard/course)
```
type RecursiveMutex struct{
    sync.Mutex
    owner   int64   //持有者
    count   int32   //重入次数
}

func (m *RecursiveMutex) Lock(){
    gid := goid.Get()
    if atomic.LoadInt(&m.owner) == gid{
        m.count++
        return
    }
    m.Mutex.Lock()
    atomic.StoreInt64(&m.owner,gid)
    m.count = 1
}
func (m *RecursiveMutex) UnLock(){
    gid := goid.Get()
    if atomic.LoadInt(&m.owner) != gid{
        panic(fmt.Sprintf("wrong the owner (%d):%d !",m.owner,gid))
    }
    m.count--
    if m.count==0{
        atomic.StoreInt64(&m.owner,-1)
        m.Mutex.UnLock()
        return
    }
}
开发中，不希望拿到goroutine id，可采用，调用者传入一个token，标示区分不同的goroutine来实现可重入锁🔒
```
###   TryLock 尝试获取排外锁
```
    const (
        mutexLocked = 1 <<iota  //加锁标识位置
        mutexWoken              //唤醒标识位子
        mutexStaring            //锁饥饿标识位置
        mutexWaiterShift = iota //标识waiter的起始bit位置
    )
    type ParkerMutex struct{
        sync.Mutex
    }
    func (m *ParkerMutex)TryLock()bool{
        //如果成功抢到锁
        if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)),0,mutexLocked){
            return true
        }
        old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
        //如果锁🔒处于唤醒、加锁或者饥饿状态下，其他goroutine准备持有，这次请求不参与了
        if old&(mutexLocked|mutexStarving|mutexWoken)!=0{
            return false
        }
        //尝试在竞争状态下请求锁
        new := old|mutexLocked
        return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)),old,new)
    }
```
###    获取等待者的数量等指标
获取Mutex结构里面的数据。
```
源码中
type Mutex struct {    
    state int32    
    sema  uint32
}
```

```
    const (
        mutexLocked = 1 <<iota  //加锁标识位置
        mutexWoken              //唤醒标识位子
        mutexStaring            //锁饥饿标识位置
        mutexWaiterShift = iota //标识waiter的起始bit位置
    )
    type ParkerMutex struct{
        sync.Mutex
    }
    func (m *ParkerMutex) Count() int {
        // 获取state字段的值
        v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
        v = v >> mutexWaiterShift //得到等待者的数值
        v = v + (v & mutexLocked) //再加上锁持有者的数量，0或者1
        return int(v)
    }
    // 锁是否被持有
    func (m *Mutex) IsLocked() bool {
        state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
        return state&mutexLocked == mutexLocked
    }
    // 是否有等待者被唤醒
    func (m *Mutex) IsWoken() bool {
        state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
        return state&mutexWoken == mutexWoken
    }
    // 锁是否处于饥饿状态
    func (m *Mutex) IsStarving() bool {
        state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
        return state&mutexStarving == mutexStarving
    }
```
###    使用Mutex实现一个线程安全队列
mutex常被用作提供一个安全操作，如队列slice，
```
type SafeSliceQueue struct{
    sync.Mutex
    data    []interface{}
}

func NewSafeSliceQueue(n int)*SafeSliceQueue{
    return &SafeSliceQueue{
        data:make([]interface{},0,n)
    }
}

func (q *SafeSliceQueue) Enqueue(x interface{}){
    q.Lock()
    q.data = append(q.data,x)
    q.UnLock()
}

func (q *SafeSliceQueue) Dequeue()inteface{}{
    q.Lock()
    if len(q.data) ==0{
        q.UnLock()
        return nil
    }
    var result = q.data[0]
    q.data = q.data[1:]
    q.UnLock()
    return result
}
```

## WaitGroup
协同等待，任务编排的一个利器


## Cond
Cond，接近弃用的一个原语

## Once
仅仅执行一次动作，常常用于单例对象的初始化场景
###  自己实现一个Once
-   有坑版
```
type ParkerOnce struct{
    done uint32
}
func (o *ParkerOnce) Do(f func()){
    if !atomic.CompareAndSwapUint32(&o.done,0,1){
        return
    }
    f() //坑点，如果f()函数执行时间过长，外部判断ParkerOnce.done已经是1完成状态，实际f里面初始化不一定已经完成
}
```
-   增强版
```
type ParkerOnce struct{
    done uint32
    m    Mutex
}
func (o *ParkerOnce) Do(f func()){
    if atomic.LoadUint32(&o.done)==0{
        o.doSlow(f)
    }
}
func (o *ParkerOnce) doSlow(f func()){
    o.m.Lock()
    defer o.m.UnLock()
    if o.done==0{
        defer atomic.StoreUint32(&o.done,1)   //    f执行完后，改标识done的值
        f()
    }
}
```
-   解决初始化函数f失败版
```
type ParkerOnce struct{
    done uint32
    m    Mutex
}
func (o *ParkerOnce) Do(f func() error) error{
    if atomic.LoadUint32(&o.done)==1{
        return nil
    }
    return  o.doSlow(f)
}
func (o *ParkerOnce) doSlow(f func() error) error{
    o.m.Lock()
    defer o.m.UnLock()
    var err error
    if o.done==0{
        err = f()
        if err == nil{
            atomic.StoreUint32(&o.done,1)   // 确认f() 是成功执行完，之后改标识done的值
        }
    }
    return err
}
```
## Map 哈希结构
### key类型，必须是可比较的

go语言中，bool、浮点、整数、复数、字符串、channel、接口，包括****包含可比较元素的struct和数组****，都是可比较的。
****而slice、map、函数值都是不可比较的。****

### map[key] 返回值
返回值可以是一个值value，也可以是两个值value,ok

### key是无序的
可采用辅助的数据结构。 如[orderedMap](https://github.com/elliotchance/orderedmap)

### 扩展map，支持并发读写
-   方法一：加锁
```
type ParkerRWMap struct{
    sync.RWMutex
    m   map[int]int
}

func NewParkerRWMap(n int) *ParkerRWMap{
    return &ParkerRWMap{
        m: make(map[int]int,n)
    }
}

func (m *ParkerRWMap)Set(k int,v int) {
    m.Lock()
    defer m.UnLock()
    m.m[k]=v
}

func (m *ParkerRWMap)Get(k int) (int,bool){
    m.Lock()
    defer m.UnLock()
    v,existed:=m.m[k]
    return v,existed
}

func (m *ParkerRWMap) Delete(k int) {
    m.Lock()
    defer m.UnLock()
    delete(m.m,k)
}
func (m *ParkerRWMap) Each(f func(k,v int)bool) {
    m.Lock()
    defer m.UnLock()
    m.m.forEach((k,v)=>{
        if !f(k,v){
            return
        }
    })
}
```
-   方法二（优化）：分片加锁
尽量减少锁的粒度和锁持有的时间，来优化全局加锁带来的性能问题。

Go中知名的项目[concurrent-map](https://github.com/orcaman/concurrent-map),默认采用32个分片
```
var SHARED_COUNT = 32
type ConcurrentMap []*ConcurrentMapShared

type ConcurrentMapShared struct{
    items   map[string]interface{}
    sync.RWMutex
}

func New() ConcurrentMap{
    m := make(ConcurrentMap,SHARED_COUNT)
    for i:=0;i<SHARED_COUNT;i++{
        m[i] = &ConcurrentMapShared{items:make(map[string]interface{})}
    }
    return m
}
//根据key计算分片索引
func (m ConcurrentMap) GetShard(key string)*ConcurrentMapShard{
    return m[uint(fnv32(key))%uint(SHARED_COUNT)]
}
```
****操作的时候，先计算分片****
```
func (m ConcurrentMap) Set(key string,val interface{}){
    shard := m.GetShard(key)
    shard.Lock()
    defer shard.UnLock()
    shard.item[key] = value
}

func (m ConcurrentMap) Get(key string)*ConcurrentMapShard{
    shard := m.GetShard(key)
    shard.RLock()
    defer shard.RUnLock()
    val,ok := shard.item[key]
    return val,ok
}
```

-   方法三：官方1.9提供的sync.Map
    -   适用场景1 : 只会增长的缓存系统中，一个 key 只写入一次而被读很多次；
    -   适用场景2 : 多个 goroutine 为不相交的键集读、写和重写键值对。
****官方给的这个适用场景不多，需评估自己业务是否符合上述特性，再适用****

## Pool，性能提升大杀器
sync.Pool，我们使用它可以创建池化的对象。（一般做性能优化的时候，会考虑采用对象池，把不用的对象回收起来，避免被gc回收掉，再使用的时候就不必在堆上重新创建了。还有如数据库连接、tcp连接等耗时操作）

### sync.Pool
保存一组可独立访问的****临时****对象.(临时意味着：它池化的对象可能会被垃圾回收掉。这对于数据库长连接等场景是不合适的)<br>
池化的对象会在未来的某个时候被毫无预兆地移除掉。而且，如果没有别的对象引用这个被移除的对象的话，这个被移除的对象就会被垃圾回收掉。

Go 内部库也用到了 sync.Pool
-   如 fmt 包，它会使用一个动态大小的 buffer 池做输出缓存，当大量的 goroutine 并发输出的时候，就会创建比较多的 buffer
****注意点****
-   1.sync.Pool本身就是线程安全的，多个goroutine可并发调用
-   2.sync.Pool不可在使用之后再复制使用

****Go 对 Pool 的优化就是避免使用锁，同时将加锁的 queue 改成 lock-free 的 queue 的实现，给即将移除的元素再多一次“复活”的机会。****
####使用方法
1.New
2.Get
3.Put
### sync.Pool的坑
#### 内存泄漏
比如初始化一个buffer池，如果往这个buffer里面增加大量数据，导致底层slice容量很大，所占的空间依然很大。而且应为Pool回收的机制，这些大的buffer可能不会被回收，而是一直占着很大空间。 造成内存泄漏问题。
#### 内存浪费
还有种就是，池子中的buffer比较大，但在使用中都只需要小的buffer，浪费空间。
#### 方案：将buffer池分成几层。 如小于1K byte大小的占一个池子，小于4K byte的元素占一个池子。
-   开源库[bucketpoll](https://github.com/vitessio/vitess/blob/master/go/bucketpool/bucketpool.go) ，提供了，你指定池子的最大和最小尺寸，可以帮你算出合适的池子数。
#### 其他方案
-   [bytebufferpool](https://github.com/valyala/bytebufferpool) 检测最大的 buffer，超过最大尺寸的 buffer，就会被丢弃。
-   [oxtoacart/bpool](https://github.com/oxtoacart/bpool) 
    -   bpool.BufferPool： 提供一个固定元素数量的 buffer 池，元素类型是 bytes.Buffer
    -   bpool.BytesPool： 提供一个固定元素数量的 byte slice 池，元素类型是 byte slice
    -   bpool.SizedBufferPool： 固定元素数量的 buffer 池
### 连接池
很常用的一个场景就是保持 TCP 的连接,而事实上，我们很少会使用 sync.Pool 去池化连接对象
#### 标准库中的 http client 池
http.Client 实现连接池的代码是在 Transport 类型中，它使用 idleConn 保存持久化的可重用的长连接
### TCP连接池
fatih 开发的[fatih/pool](https://github.com/fatih/pool) 可管理的是更通用的 net.Conn，不局限于 TCP 连接
### 数据库连接池
标准库 sql.DB 数据库的连接池。默认的 MaxIdleConns 是 2，这太小了。
<br>DB 的 freeConn 保存了 idle 的连接，这样，当我们获取数据库连接的时候，它就会优先尝试从 freeConn 获取已有的连接
### Memcached Client 连接池
https://time.geekbang.org/column/article/301716
### Worker Pool
**last but not least!!!**
<br>应用得非常广泛的场景.创建一个 Worker Pool 来减少 goroutine 的使用
如：fasthttp用来处理TCP连接的 [workerPool](https://github.com/valyala/fasthttp/blob/9f11af296864153ee45341d3f2fe0f5178fd6210/workerpool.go#L16)

-   gammazero/workerpool：
    -   gammazero/workerpool 可以无限制地提交任务，提供了更便利的 Submit 和 SubmitWait 方法提交任务，还可以提供当前的 worker 数和任务数以及关闭 Pool 的功能。
-   ivpusic/grpool：
    -   grpool 创建 Pool 的时候需要提供 Worker 的数量和等待执行的任务的最大数量，任务的提交是直接往 Channel 放入任务。
-   dpaks/goworkers：
    -   dpaks/goworkers 提供了更便利的 Submi 方法提交任务以及 Worker 数、任务数等查询方法、关闭 Pool 的方法。它的任务的执行结果需要在 ResultChan 和 ErrChan 中去获取，没有提供阻塞的方法，但是它可以在初始化的时候设置 Worker 的数量和任务数。
    
## Context
-   上下文（Context）
-   超时（Timeout）和取消（Cancel）的机制      (context命名--“名不副实”)
```
type Context interface {    
    Deadline() (deadline time.Time, ok bool)    //返回这个 Context 被取消的截止日期
    Done() <-chan struct{}      //返回一个 Channel 对象
    Err() error    
    Value(key interface{}) interface{}   //返回此 ctx 中和指定的 key 相关联的 value
}
```
### 特殊用途 Context 的方法
#### WithValue
#### WithCancel
#### WithTimeout
#### WithDeadline

## atomic 原子操作
**因为不同的 CPU 架构甚至不同的版本提供的原子操作的指令是不同的，所以，要用一种编程语言实现支持不同架构的原子操作是相当有难度的。**
### atomic 原子操作的应用场景


