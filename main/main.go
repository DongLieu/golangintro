package main

import (
	// "crypto/sha256"
	"fmt"

	"choiroixoangay/blockchain"
)

//
// type Block struct {
// 	Hash     []byte
// 	Data     []byte
// 	PrevHash []byte
// }
//
// type BlockChain struct {
// 	blocks []*Block
// }
//
// func (block *Block) BamHash() {
// 	giatriBam := bytes.Join([][]byte{block.Data, block.PrevHash}, []byte{})
// 	hash := sha256.Sum256(giatriBam)
// 	block.Hash = hash[:]
//
// }
//
// func CreateBlock(data string, PrevHash []byte) *Block {
// 	block := &Block{[]byte{}, []byte(data), PrevHash}
// 	block.BamHash()
// 	return block
// }
//
// func Genesis() *Block {
// 	return CreateBlock("Genesis", []byte{})
// }
//
// func khoitaoBlockChain() *BlockChain {
// 	return &BlockChain{[]*Block{Genesis()}}
// }
//
// func (chain *BlockChain) AddBlock(data string) {
// 	prevBlock := chain.blocks[len(chain.blocks)-1]
// 	block_new := CreateBlock(data, prevBlock.Hash)
// 	chain.blocks = append(chain.blocks, block_new)
// }

func main() {
	noihello := blockchain.SayHello()
	fmt.Println(noihello)

	chain := blockchain.khoitaoBlockChain()
	chain.AddBlock("Alice gui Bob 10$")
	for _, block := range chain.blocks {
		fmt.Printf("	Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("	Data in Block: %s\n", block.Data)
		fmt.Printf("	Hash: %x\n\n", block.Hash)
	}
}
