package main

// 1.引入http包
// 2.定义路由匹配模式 与对应的处理函数
// 3.创建http server对象指定监听端口、handler 处理器
// 4.调用server的ListenAndServe()函数

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

	//  curl -X POST http://127.0.0.1:8088/demo/foo -H 'Content-Type: application/json' -d '{"login":"login","password":"password"}'
	http.HandleFunc("/demo/foo", handFooFunc)

	//http://127.0.0.1:8088/toolkit.html
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
