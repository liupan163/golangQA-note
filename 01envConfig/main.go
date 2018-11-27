package envConfig

func main() {
}
//配置相关环境
//GOROOT、GOPATH和GOBIN

//一、设置GOPATH有什么意义？

//GOROOT其实就是golang的安装路径

环境变量GOPATH的值可以是一个目录的路径，也可以包含多个目录路径，每个目录都代表Go语言的一个工作区。
这些工作区用于放置Go语言的源码文件，以及安装后的归档文件（archive file，就是以.a为扩展名的文件）和可执行文件。

pkg子目录---归档文件；
bin子目录---可执行文件；
src---源码路径。



