package _2sourceCode

/*import "flag"

//go语言标准库中---flag代码包---专门用于接受和解析命令参数。

flag.StringVar(&name,"name","everyone","The greeting object")

flag.StringVar接受4个参数，第1个用于存储该命令参数值得地址。变量name的地址，&name
第2个参数为该命令指定命令参数名称，name。
第3个指定默认值
第4只作为简短说明，打印命令说明时会用到。

区别于flag.String。  后者会返回个分配好存储命令参数地址。
改动var name string ---> var name = flag.String("name","everyone","The greeting object")

flag.Parse()用于真正解析命令参数，并把他们的值赋给相应的变量。  顺序，声明和设置之后

查看命令源码文件的参考说明：
go run demo.go --help

自定义命令源码的参数使用说明：
flag.Usage重新赋值。

不用全局的flag.CommandLine变量，自己创建一个私有命令参数容器。
var cmdLine = flag.NewFlagSet("question",flag.ExitOnError)
替换成cmdLine.Parse(os.Args[1:])
其中os.Args[1:]指的就是我们给定的那些命令参数。
*/
