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

func init() {
	flag.Usage()
	flag.StringVar(&baseAddress, "host", "-1", "用于指定端口")
	flag.IntVar(&start, "start", 21, "用于指定开始范围，默认为21")
	flag.IntVar(&end, "end", 200, "用于指定端口结束范围，默认为200")
}

func main() {
	flag.Parse()

	if baseAddress == "-1" {
		fmt.Println("请指定主机地址")
		return
	}

	/*等待组*/
	var waitGroup sync.WaitGroup
	startTime := time.Now()
	for i := start; i <= end; i++ {
		address := fmt.Sprintf(baseAddress+":%d", i)
		waitGroup.Add(1)
		go func(j int, w *sync.WaitGroup) {
			defer w.Done()
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%d 关闭了\n", j)
			} else {
				err = conn.Close()
				fmt.Printf("%d 打开了\n", j)
			}
		}(i, &waitGroup)
	}
	waitGroup.Wait()
	elspased := time.Since(startTime) / 1e9
	fmt.Printf("\n\n%d seconds", elspased)
}
