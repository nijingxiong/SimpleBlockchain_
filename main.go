package main

//import "fmt"

func main() {
	//PART 2:
	//bc := NewBlockchain()
	//
	//bc.AddBlock("Send 1 BTC to Ivan")
	//bc.AddBlock("Send 2 more BTC to Ivan")
	//
	//for _, block := range bc.blocks {
	//	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	pow:=NewProofWork(block)
	//	fmt.Printf("PoW:%s\n",strconv.FormatBool(pow.Validate()))
	//	fmt.Println()
	//测试
		//num:=int64(256)
		//buff:=new(bytes.Buffer)//Buffer是一个实现了读写方法的可变大小的字节缓冲
		//err:=binary.Write(buff,binary.BigEndian,num)
		//if err==nil{
		//	fmt.Println(buff.Bytes())
		//}
		//data:=bytes.Join([][]byte{buff.Bytes(),[]byte("hello")},[]byte{35})
		//fmt.Println(data)
		//fmt.Println([]byte("hello"))
	bc:=NewBlockchain()
	defer bc.db.Close()
	cli:=CLI{bc}
	cli.Run()
	}

