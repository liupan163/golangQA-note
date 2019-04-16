package httpInfo

import (
	"net/http"
	"fmt"
	"os"
)

/*
	https连接过程:
    1、客户端发送请求到服务器端
    2、服务器端返回证书和公开密钥，公开密钥作为证书的一部分而存在
    3、客户端验证证书和公开密钥的有效性(TLS部分来做的)，如果有效，则生成 共享密钥（随机值） 并使用 公开密钥 加密发送到服务器端
    4、服务器端使用私有密钥解密数据，并使用收到的共享密钥（随机值）加密数据，发送到客户端
    5、客户端使用共享密钥解密数据
    6、SSL加密建立………
*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, this is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	_, err := os.Open("cert_server/server.crt")
	if err != nil {
		panic(err)
	}
	http.ListenAndServeTLS(":8081", "cert_server/server.crt",
		"cert_server/server.key", nil)
}
