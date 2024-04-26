package main

import (
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func handFooFunc(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
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
	prometheus.MustRegister(pingCounter)

	http.Handle("/metrics", promhttp.Handler())
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

	myserver.ListenAndServe()

}
