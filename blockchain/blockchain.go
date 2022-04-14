package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevhash"`
	Height   int    `json:"height"`
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totlaBlocks := len(GetBlockchain().blocks)
	if totlaBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totlaBlocks-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{Data: data, Hash: "", PrevHash: getLastHash(), Height: len(b.AllBlocks()) + 1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis block")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}

func (b *blockchain) GetBlock(height int) (*block, error) {
	if height > len(b.blocks) {
		return nil, errors.New("block not found")
	}
	return b.blocks[height-1], nil
}
