package main

import (
	"fmt"
	"os"
	"flag"
	"log"
)

type CLI struct {
	bc *BlockChain
}
const usage =`
Usage:
	addblock -data BLCOK_DATA add a block to the blockchain
	printchain                print all the blocks of the blockchain
`
func (cli *CLI) printUsage() {
	fmt.Println(usage)
}
//验证命令行的合法性
func (cli *CLI) validateArgs(){
	if len(os.Args)<2{
		cli.printUsage()
		os.Exit(1)
	}
}
func (cli *CLI) Run()  {
	cli.validateArgs()
	//使用flag包来解析命令行参数
	addBlockCmd:=flag.NewFlagSet("addblock",flag.ExitOnError)
	printChainCmd:=flag.NewFlagSet("printchain",flag.ExitOnError)
	addBlockData:=addBlockCmd.String("data","","Block data")
	switch os.Args[1] {
	case "addblock":
		err:=addBlockCmd.Parse(os.Args[2:])
		if err!=nil{
			log.Panic(err)
		}
	case "printchain":
		err:=printChainCmd.Parse(os.Args[2:])
		if err!=nil{
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed(){
		if *addBlockData==""{
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.bc.AddBlock(*addBlockData)
	}
	if printChainCmd.Parsed(){
		cli.printChain()
	}
}