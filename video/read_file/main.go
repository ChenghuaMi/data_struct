package main

import (
	"fmt"
	"io/ioutil"
)

// 实现 并发的 读取 文件
// 读取 file1.txt,file2.txt, file3.txt

func readFile(ch chan string, filename string) {
	bt, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	ch <- string(bt)
}
func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	ch := make(chan string)
	for _, file := range files { //并发的读取文件
		go readFile(ch, file)
	}
	//读取管道中的文件内容
	for i := 0; i < len(files); i++ {
		fmt.Println(<-ch)
	}
}
