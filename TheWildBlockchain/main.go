package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//区块
type blockmain struct {
	whattime      string //时间戳
	LastBlockHash string //父哈希
	BlockHash     string //当前哈希
	Index         int    //当前位置
	date          string //交易内容
	BPM           int    //生成区块的频率
}

//区块链，一个区块的slice。
var Blockchain []blockmain

//用父哈希等算出来当前哈希
//代码来自gobyexample.com
func GetBlockHash(b blockmain) (hashnow string) {
	t := time.Now()
	check := string(b.Index) + b.date + b.LastBlockHash + t.String()
	h := sha256.New()
	h.Write([]byte(check))
	return hex.EncodeToString(h.Sum(nil)) //返回16进制编码。
}

//用父哈希等算出来当前哈希。
//用blockmain struct解决问题，分开处理新旧blockmain，输入旧blockmain，返回新的blockmain。
func OldAndNew(oldBlock blockmain, BPM int) (blockmain, error) {
	var newBlock blockmain
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1 //区块数加一
	newBlock.whattime = t.String()
	newBlock.BPM = BPM                          //频率
	newBlock.LastBlockHash = oldBlock.BlockHash //这里的老hash来自新hash。
	newBlock.BlockHash = GetBlockHash(newBlock)
	return newBlock, nil
}

//块验证
func blockcheck(newBlock, oldBlock blockmain) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.BlockHash != newBlock.LastBlockHash {
		return false
	}

	if GetBlockHash(newBlock) != newBlock.BlockHash {
		return false
	}

	return true
}

//一个验证函数，选择更长的链作为真链，输入值newBlocks为区块链的链
func replaceChain(newBlocks []blockmain) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
