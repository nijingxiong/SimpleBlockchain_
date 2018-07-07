package main

import (
	"bytes"
	"encoding/binary"//binary包实现了简单的数字与字节序列的转换以及变长值的编解码
	"log"
)

//将一个int64转化成一个字节数组（byte array）
func IntToHex(num int64)[]byte  {
	buff:=new(bytes.Buffer)//Buffer是一个实现了读写方法的可变大小的字节缓冲
	err:=binary.Write(buff,binary.BigEndian,num)//将num的binary编码格式写入buff
	if err!=nil{
		log.Panic(err)
	}
	return buff.Bytes()//返回[]byte，返回未读取部分字节数据的切片
}
