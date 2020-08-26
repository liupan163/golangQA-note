# package channel

无缓冲 Chan 的发送和接收是否同步?
ch := make(chan int)    无缓冲的channel由于没有缓冲发送和接收需要同步.
ch := make(chan int, 2) 有缓冲channel不要求发送和接收操作同步. 
channel无缓冲时，发送阻塞直到数据被接收，接收阻塞直到读到数据。
channel有缓冲时，当缓冲满时发送阻塞，当缓冲空时接收阻塞。


