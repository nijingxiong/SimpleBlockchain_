package main

type BlockChain struct {
	blocks []*Block //block数组
}
func (bc *BlockChain) AddBlock(data string){
	prevBlock:=bc.blocks[len(bc.blocks)-1]
	newBlock:=NewBlock(data,prevBlock.Hash)
	bc.blocks=append(bc.blocks,newBlock)
}
func NewGenesisBlock() *Block{
	return NewBlock("Genesis Block",[]byte{})
}
func NewBlockchain() *BlockChain  {
	return &BlockChain{[]*Block{NewGenesisBlock()}}

}
