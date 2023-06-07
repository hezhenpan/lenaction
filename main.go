package main

import (
	"github.com/hezhenpan/aws-private-lib/awpkg"
	"time"
)

func main() {
	plus := awpkg.GetNBPlus("zhangsan", "lisi")
	v := time.Now().String()
	println("my time version", v)
	println("hello world", plus)
}
