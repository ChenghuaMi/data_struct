package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

//深拷贝核心逻辑：将源对象序列化（编码）为字节流，再将字节流反序列化（解码）为新对象，
//新对象与原对象内存完全隔离，修改互不影响。

type Address struct {
	City  string
	Phone []string
}
type User struct {
	Id      int
	Name    string
	Address *Address
}

func DeepCopy(dst, src interface{}) error {

	buff := new(bytes.Buffer) //创建缓冲区，用于存储序列化后的字节流

	// 1. 编码（序列化）源对象到缓冲区
	encoder := gob.NewEncoder(buff)
	if err := encoder.Encode(src); err != nil {
		fmt.Println("encode err:", err)
		return err
	}
	//2. 解码（反序列化）缓冲区数据到目标对象
	decoder := gob.NewDecoder(buff)
	if err := decoder.Decode(dst); err != nil {
		fmt.Println("decode err:", err)
		return err
	}
	return nil
}
func main() {
	var dst User
	src := User{
		Id:   90,
		Name: "zs",
		Address: &Address{
			City:  "sz",
			Phone: []string{"123"},
		},
	}
	DeepCopy(&dst, &src)

	dst.Id = 10
	fmt.Printf("%#v\n", dst)
	fmt.Printf("%#v\n", src)
}
