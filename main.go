package main

/**
https://gowebexamples.com/
1. 引入 http 包
2. 定义路由匹配模式与对应的处理函数
3. 创建 http server对象，指定监听端口和 handler 处理器类型
4. 调用server的ListenAndServe()函数
*/
import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func handFooFunc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {

	// curl -X POST http://127.0.0.1:8088/demo/foo -H 'Content-Type: application/json' -d '{"login":"login","password":"password"}'
	http.HandleFunc("/demo/foo", handFooFunc)

	// http://127.0.0.1:8088/toolkit.html
	http.Handle("/", http.FileServer(http.Dir("static/")))

	myserver := &http.Server{
		Addr:           ":8088",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    180 * time.Second,
		WriteTimeout:   180 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("server serving ...")
	myserver.ListenAndServe()

}
