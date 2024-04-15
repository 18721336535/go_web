package main

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
	//  curl -X POST http://127.0.0.1:8088/demofoo -H 'Content-Type: application/json' -d '{"login":"login","password":"password"}'
	http.HandleFunc("/demo/foo", handFooFunc)

	//http://127.0.0.1:8088/toolkit.html
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", fs)

	myserver := &http.Server{
		Addr:           ":8088",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    180 * time.Second,
		WriteTimeout:   180 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	myserver.ListenAndServe()

}
