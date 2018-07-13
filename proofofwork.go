package main

import (
	"math/big"
	"bytes"

	"fmt"
	"math"
	"crypto/sha256"
)

const  targetBits  =24 //挖矿难度值,指算出来的哈希前24位必须是0
const maxNonce=math.MaxInt64

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofWork(b *Block) *ProofOfWork  {
	target:=big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))//将target左移256-targetBits位
	pow:=&ProofOfWork{b,target}//将新生成的hash与pow.target对比，当比target小时是一个有效的证明
	return pow
}
//为生成hash准备数据，数据由上一个块的hash，块数据时间戳和挖矿难度和nonce组成
func (pow *ProofOfWork) prepareData(nonce int)[]byte  {
	//将下列字节数组以空为连接词连成一个字节数组
	data:=bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),//将timestamp转换为字节字节数组
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	},
	[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run()(int,[]byte){
	var hashInt big.Int//hash的整型表示
	var hash [32]byte
	nonce:=0

	fmt.Println("Mining the block containing%s\n",string(pow.block.Data[:]))
	for nonce<maxNonce{//避免nonce溢出
		data:=pow.prepareData(nonce)
		hash=sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target)==-1{
			fmt.Printf("\r%x",hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce,hash[:]
}
//对工作量证明进行验证
func (pow *ProofOfWork) Validate() bool{
	var hashInt big.Int
	data:=pow.prepareData(pow.block.Nonce)
	hash:=sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isVaild:=hashInt.Cmp(pow.target)==-1
	return isVaild
}