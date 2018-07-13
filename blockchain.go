package main

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbFile="blockchain.db"
const blockBucket="blocks"
/*tip用来存储最后一个块的哈希
在链的末端可能出现短暂分叉的情况，选择tip也就是选择了哪条链
db用来存储数据库连接
 */
type BlockChain struct {
	tip []byte
	db *bolt.DB
}
//加入区块时，需要将区块持久化到数据库中
func (bc *BlockChain) AddBlock(data string){
	var lastHash []byte
	//获取最后一个块的哈希用于生成新块的哈希
	err:=bc.db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		lastHash=b.Get([]byte("1"))
		return nil
	})
	if err!=nil{
		log.Fatal(err)
	}
	//生成新的块并添加进数据库中
	newBlock:=NewBlock(data,lastHash)
	err=bc.db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		err:=b.Put(newBlock.Hash,newBlock.Serialize())
		if err!=nil{
			log.Fatal(err)
		}
		err=b.Put([]byte("1"),newBlock.Hash)
		bc.tip=newBlock.Hash
		return nil
	})
}
func NewGenesisBlock() *Block{
	return NewBlock("Genesis Block",[]byte{})
}
func NewBlockchain() *BlockChain  {
	var tip []byte
	db,err:=bolt.Open(dbFile,0600,nil)//打开boltdb文件的标准做法，不存在就创建一个这个文件

	err=db.Update(func(tx *bolt.Tx)error{//打开一个读写事务
	//获取存储区块的bucket
		b:=tx.Bucket([]byte(blockBucket))
		//如果数据库中不存在区块链就创建一个，否则直接读取最后一个块的哈希
		if b==nil{
			genesis:=NewGenesisBlock()
			b,_:=tx.CreateBucket([]byte(blockBucket))
			err=b.Put(genesis.Hash,genesis.Serialize())
			err=b.Put([]byte("1"),genesis.Hash)
			tip=genesis.Hash
		}else {
			tip=b.Get([]byte("1"))
		}
		return nil
	})
	bc:=BlockChain{tip:tip,db:db}
	return &bc
}
//区块链迭代器
type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}
func (bc *BlockChain)Iterator() *BlockchainIterator{
	bci:=&BlockchainIterator{bc.tip,bc.db}
	return bci
}
//返回链中的下一个块
func (i *BlockchainIterator) Next() *Block{
	var block *Block
	err:=i.db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		encodedBlock:=b.Get(i.currentHash)
		block=DeserializeBlock(encodedBlock)
		return nil
	})
	if err!=nil{
		log.Panic(err)
	}
	i.currentHash=block.PrevBlockHash
	return block
}