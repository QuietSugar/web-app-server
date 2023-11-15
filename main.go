package main

import (
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"time"
)

var (
	root = flag.String("w", ".", "Root dir of the server")
)

func init() {
	flag.Parse()
}

var url string

func main() {

	a, err := filepath.Abs(*root)
	if err != nil {
		log.Fatal(err)
	}
	d := http.Dir(a)
	fs := http.FileServer(d)

	// 创建一个新的多路复用器
	mux := http.NewServeMux()

	// 设置文件服务器处理器
	mux.Handle("/", fs)

	// 获取一个可用的随机端口
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	// 获取实际监听的端口
	port := listener.Addr().(*net.TCPAddr).Port
	fmt.Printf("文件服务器已启动，访问地址：http://127.0.0.1:%d\n", port)
	go func() {
		// 启动服务器
		err = http.Serve(listener, mux)
		if err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(time.Second)
	url = fmt.Sprintf("http://localhost:%d", port)
	open.Run(url)
	startTray()
	fmt.Println("HTTP 服务器已在随机端口", port, "上启动")
}
