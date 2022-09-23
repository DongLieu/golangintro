package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/DongLieu/golangintro/blockchain"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 { // neu k them print hoac add thi se hien thi huong dan va thoat
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator() // tra ve lash hash va db o moi thoi diem

	for {
		block := iter.Next() // thoi diem sau

		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("pre.Hash: %x\n", block.PrevHash)
		pow := blockchain.NewProof(block) // PoW :block, tager
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data") //tra ve dia chi string data, neu ko co mac dinh la ""
	switch os.Args[1] {                                           // os.Args[1] = add, print
	case "add":
		err := addBlockCmd.Parse(os.Args[2:]) // os.Args[2:] = -block data
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() { //true neu Parse da duoc goi
		if *addBlockData == "" { //neu ko co data thi lai defau
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}
	if printChainCmd.Parsed() { //true neu Parse da duoc goi
		cli.printChain()
	}
}

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockChain() //lasthask+db
	defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.run()

}
