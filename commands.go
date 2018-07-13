package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) addBlock(data string){
	cli.bc.AddBlock(data)
	fmt.Println("success")
}
func (cli *CLI) printChain()  {
	bci:=cli.bc.Iterator()
	for{
		block:=bci.Next()
		fmt.Printf("Prev hash: %x \n",block.PrevBlockHash)
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Printf("Hash: %x\n",block.Hash)
		pow:=NewProofWork(block)
		fmt.Printf("PoW: %s\n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.PrevBlockHash)==0{
			break
		}
	}
}