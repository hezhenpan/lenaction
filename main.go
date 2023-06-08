package main

import (
	"github.com/hezhenpan/aws-private-lib/awpkg"
	"time"
)

func main() {
	for {
		plus := awpkg.GetNBPlus("zhangsan", "lisi")
		v := time.Now().String()
		println("my time version", v)
		println("hello world", plus)
		println("sleep 10s.")
		time.Sleep(10 * time.Second)
	}
}
