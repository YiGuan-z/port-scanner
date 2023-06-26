package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

/*扫描地址*/
var baseAddress string

/*开始范围*/
var start int

/*结束范围*/
var end int

/*缓冲*/
var cache int

func init() {
	flag.Usage()
	flag.StringVar(&baseAddress, "host", "-1", "用于指定端口")
	flag.IntVar(&start, "start", 21, "用于指定开始范围，默认为21")
	flag.IntVar(&end, "end", 200, "用于指定端口结束范围，默认为200")
	flag.IntVar(&cache, "cache", 100, "用于指定发送速率，默认为100")
}

func main() {
	flag.Parse()

	if baseAddress == "-1" {
		fmt.Println("请指定主机地址")
		return
	}

	ports := make(chan int, cache)
	/*等待组*/
	var waitGroup sync.WaitGroup
	startTime := time.Now()

	for i := 0; i <= cache; i++ {
		go worker(ports, &waitGroup)
	}

	for i := start; i <= end; i++ {
		waitGroup.Add(1)
		ports <- i
	}

	waitGroup.Wait()
	close(ports)
	elspased := time.Since(startTime) / 1e9
	fmt.Printf("\n\n%d seconds🫡", elspased)
}
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", baseAddress, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%d 端口已关闭🤡\n", p)
		} else {
			_ = conn.Close()
			fmt.Printf("%d 端口已开启😁\n", p)
		}
		wg.Done()
	}
}
