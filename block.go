package main

import (
	"time"
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