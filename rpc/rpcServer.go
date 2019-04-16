package rpc

import (
	"errors"
	"net/rpc"
	"net"
	"log"
	"fmt"
	"os"
	"net/http"
)

//rpc框架负责屏蔽底层的传输过程、序列化方式、通信细节
//参考资料 https://studygolang.com/articles/14336

//     1 net/rpc库
//     2 net/rpc/jsonrpc库
//     3 protorpc库
func main() {
	//mainNetRpcS()
	mainJsonRpcS()
	//mainPbRpcS()
}

type Arith struct{}
type ArithRequest struct {
	Req1 int
	Req2 int
}
type ArithResponse struct {
	Multiply  int
	Divide    int
	Remainder int
}

func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Multiply = req.Req1 * req.Req2
	return nil
}
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.Req2 == 0 {
		return errors.New("divide by zero")
	}
	res.Divide = req.Req1 / req.Req2
	res.Remainder = req.Req1 % req.Req2
	return nil
}

//eg1
func mainNetRpcS() {
	rpc.Register(new(Arith)) //注册rpc服务
	rpc.HandleHTTP()         //采用http协议作为rpc载体
	lis, err := net.Listen("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("server fatal err", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection by rpc")
	http.Serve(lis, nil)
}

//eg2
func mainJsonRpcS() {
	rpc.Register(new(Arith))
	lis, err := net.Listen("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("server fatal err", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection by jsonrcp")
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
		}(conn)
	}
}

//eg3
//protorpc
/*我们需要根据上述定义的arith.proto文件生成RPC服务代码。
	要先安装protorpc库：go get github.com/chai2010/protorpc
	然后使用protoc工具生成代码：protoc --go_out=plugin=protorpc=. arith.proto
	执行protoc命令后，在与arith.proto文件同级的目录下生成了一个arith.pb.go文件，里面包含了RPC方法定义和服务注册的代码。
	基于生成的arith.pb.go代码我们来实现一个rpc服务端
*/
/*
func (this *Arith) Multiply(req *pb.ArithRequest, res *pb.ArithRequest) error {
	res.Pro = req.getA() * req.GetB()
	return nil
}
func (this *Arith) Divide(req *pb.ArithRequest, res *pb.ArithRequest) error {
	if req.getB() == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.getA() / req.getB()
	res.Rem = req.getA() % req.getB()
	return nil
}
func mainPbRpcS() {
	pb.ListenAndSerArithService("tcp", "127.0.0.1:8097", new(Arith))
}
*/
