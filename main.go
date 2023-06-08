package main

import (
	"fmt"
	"github.com/hezhenpan/aws-private-lib/awpkg"
	"time"
)

func main() {
	for {
		plus := awpkg.GetNBPlus("zhangsan", "lisi")
		v := time.Now().String()
		fmt.Println("my time version", v)
		fmt.Println("hello world", plus)
		fmt.Println("sleep 10s.")
		time.Sleep(10 * time.Second)
	}
}
