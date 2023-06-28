package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	/*扫描地址*/
	baseAddress string
	/*开始范围*/
	start int
	/*结束范围*/
	end int
	/*缓冲*/
	cache int
)

func init() {
	flag.StringVar(&baseAddress, "host", "-1", "用于指定端口")
	flag.IntVar(&start, "start", 1, "用于指定开始范围，默认为1")
	flag.IntVar(&end, "end", 500, "用于指定端口结束范围，默认为500")
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

// worker todo 需要一个扫描端口的逻辑
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", baseAddress, p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			fmt.Printf("%d 端口已关闭🤡\n", p)
		} else {
			fmt.Printf("%d 端口已开启😁\n", p)
			_ = conn.Close()
		}
		wg.Done()
	}
}
