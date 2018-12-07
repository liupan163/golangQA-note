//条件变量sync.Cond

//条件变量是基于互斥锁的，它必须有互斥锁才能发挥作用。

等待通知wait、单发通知signal、广播通知broadcast

条件变量怎么跟互斥锁配合使用？

var mailbox uint8
var lock sync.RWMutex
sendCond := sync.NewCond(&lock)
recvCond := sync.NewCond(lock.RLocker())
lock.Lock()
for mailbox == 1 {
 sendCond.Wait()
}
mailbox = 1
lock.Unlock()
recvCond.Signal()

//*sync.Cond类型的值可以被传递吗？那 sync.Cond类型的值呢？
