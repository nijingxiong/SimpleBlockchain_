package main

import (
	"time"
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Timestamp int64 //当前时间戳
	Data []byte //区块存储的有效信息
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}

//func (b *Block) SetHash()  {
//	timestamp:=[]byte(strconv.FormatInt(b.Timestamp,10))
//	headers :=bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestamp},[]byte{})//以{}为连接符将前三个字符串类型连接成一个字符串
//	hash:=sha256.Sum256(headers)
//	b.Hash=hash[:]
//
//}
func NewBlock(data string,prevBlockHash []byte) *Block  {
	block :=&Block{time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{},0}
	pow:=NewProofWork(block)
	nonce,hash:=pow.Run()
	block.Hash=hash[:]
	block.Nonce=nonce

	return block

}
//序列化
func (b *Block)Serialize() []byte{
	var result bytes.Buffer//定义一个buffer用来存储序列化后的数据
	encoder:=gob.NewEncoder(&result)//
	err:=encoder.Encode(b)//对block进行编码并写入result中
	if err!=nil{
		log.Fatal(err)
	}
	return result.Bytes()
}
//反序列化
func DeserializeBlock(d []byte) *Block{
	var block Block
	decoder:=gob.NewDecoder(bytes.NewReader(d))
	err:=decoder.Decode(&block)
	if err!=nil{
		log.Fatal(err)
	}
	return &block
	}