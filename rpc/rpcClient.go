package rpc

import (
	"net/rpc"
	"log"
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	//mainNetRpcC()
	mainJsonRpcC()
	//mainPbRpcC()
}

//eg1
type ArithReq struct {
	Req1 int
	Req2 int
}
type ArithRes struct {
	Multiply  int
	Divide    int
	Remainder int
}

func mainNetRpcC() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("client fatal conn", err)
	}
	req := ArithReq{9, 2}
	var res ArithRes

	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatalln("arith multiply err", err)
	}
	fmt.Printf("multiply res: %d*%d = %d \n", req.Req1, req.Req2, res.Multiply)

	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("arith divide err", err)
	}
	fmt.Printf("divider res:%d / %d = %d\n", req.Req1, req.Req2, res.Divide)
}

//eg2
func mainJsonRpcC() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("client conn fatal", err)
	}
	req := ArithReq{9, 2}
	var res ArithRes

	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.Req1, req.Req2, res.Multiply)

	err = conn.Call("Arith.Divide", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.Req1, req.Req2, res.Divide)
}

//eg3
/*
func mainPbRpcC() {
	conn, err := pb.DialArithService("tcp", "127.0.0.1")
	if err != nil {
		log.Fatalln("client conn fatal", err)
	}
	req := &pb.ArithRequest{9, 2}

	res, err := conn.Multiply(req)
	if err != nil {
		log.Fatalln("arith error:", err)
	}
	fmt.Printf("%d * %d=%d\n", req.GetA, req.GetB, res.GetPro())

	res, err = conn.Divide(req)
	if err != nil {
		log.Fatalln("arith error:", err)
	}
	fmt.Printf("%d / %d=%d\n", req.A, req.GetB, res.Quo, res.Rem)
}
*/
