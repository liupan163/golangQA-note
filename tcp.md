tcp

TCP 有哪些状态?
close
listen     服务端，等待连接的
syn_sent    发送连接请求
syn_received    收到连接请求后
established 建立连接请求

fin_wait_1  主动发起通信结束方，等待对方响应
fin_wait_2  主动发起通信结束方，收到对方ACK包之后
time_wait   主动发起方，最后一条ack后，会开始一段等待的时间

-   建立一个 socket 连接要经过哪些步骤?
    socket()函数
    bind()函数
    listen()、connect()函数
    accept()函数
    read()、write()函数等
    close()函数

客户端发送connect（）方法，就是tcp发送了请求
服务端调用accept（）方法，并返回，就是服务端返回tcp请求
客户端这时connect（）方法返回，就是客户端对服务端消息的返回。
服务段收到accept（）方法返回， 至此，三次握手完成

客户端的connect在三次握手的第二个次返回，而服务器端的accept在三次握手的第三次返回。

