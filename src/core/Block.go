package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int64 //区块编号
	Timestamp int64 // 区块时间戳
	PrevBlockHash string //上一个区块的哈希值
	Hash string //当前区块的哈希值

	Data string //区块数据
}

func CalculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash +b.Data
	hashBytes :=sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashBytes[:])
}

func GenerateNewBlock(prevBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index+1
	newBlock.PrevBlockHash = prevBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block{
	prevBlock := Block{}
	prevBlock.Index = -1
	prevBlock.Hash = ""
	return GenerateNewBlock(prevBlock,"Genesis Block")
}
