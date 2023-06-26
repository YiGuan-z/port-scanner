package main

import (
	"flag"
	"fmt"
	"net"
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
	for i := 21; i < 120; i++ {
		address := fmt.Sprintf(baseAddress+":%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%d 关闭了\n", i)
			continue
		}
		_ = conn.Close()
		fmt.Printf("%d 打开了\n", i)

	}
}
