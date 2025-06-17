package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	go func() {
		fmt.Println("1")
	}()
	go func() {
		fmt.Println("133")
	}()
	go func() {
		fmt.Println("322")
	}()
	defer trace.Stop()
	// 业务逻辑...
}
