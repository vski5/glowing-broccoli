package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

//区块
type blockmain struct {
	whattime      int64  //时间戳
	LastBlockHash []byte //父哈希
	BlockHash     []byte //当前哈希
	date          []byte //交易内容
}

//用父哈希等算出来当前哈希，向下运行时的扣子。
func (b *blockmain) SetBlockHash() {
	//根据当前时间生成[]byte，然后用sha256加密一下。
	t := time.Now()
	checktime := []byte(t.String())
	checktime2 := sha256.Sum256(checktime)
	checktime3 := checktime2[:]
	//算整个区块头的sha256
	checkhash := bytes.Join([][]byte{checktime3, b.date, b.LastBlockHash}, []byte{})
	BlockHash := sha256.Sum256(checkhash)
	//算出来的hash赋值给BlockHash（当前哈希）
	b.BlockHash = BlockHash[:]
}

//将交易内容（date）和LastBlockHash写入当前的区块。
//也就是生成新区块
func SetNowBlock(date string, LastBlockHash []byte) *blockmain {
	b := &blockmain{
		whattime:      time.Now().Unix(), //时间戳
		LastBlockHash: LastBlockHash,     //父哈希
		BlockHash:     []byte{},          //当前哈希
		date:          []byte(date)}      //交易内容

	b.SetBlockHash()
	return b
}

//写入第一笔交易和首个LastBlockHash，也就是创造创世区块。这里用的是创世区块生成的时间。
//这里只是方法，还没开始操作。
//开始创世还用这个方法把数据写到链上去。分两步走。
func firstblock() *blockmain {
	t := time.Now()
	firstBlockHash := []byte(t.String())
	firstBlockHash2 := sha256.Sum256(firstBlockHash)
	firstBlockHash3 := firstBlockHash2[:]
	return SetNowBlock("firstdate", firstBlockHash3)
}

//将创世区块与后续的区块连起来,实现链
//区块链的链，是区块这个struct类型的指针的数组
type BlockMainChain struct {
	blocks []*blockmain
}

/*开始创世还用这个方法把数据写到链上去
func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}
*/

//将创世区块加入这个链中，返回值写到链里面去。
func firstblockmainin() *BlockMainChain {
	return &BlockMainChain{[]*blockmain{firstblock()}}

}

//AddBlock 向链中加入一个新块
//data 在实际中就是交易
//我最看不懂的一部分。
func (Chain *BlockMainChain) AddBlock(data string) {
	//Chain.blocks是区块这个struct类型的指针的数组，对这个指针类型的数据进行修改
	//此处是算出来共有多少个区块。
	lastBlock := Chain.blocks[len(Chain.blocks)-1]
	//创造下一个区块，自己取代自己。
	SetNowBlock := SetNowBlock(data, lastBlock.BlockHash)
	//将上面创造的区块SetNowBlock加到Chain.blocks（blocks []*blockmain）里面。
	Chain.blocks = append(Chain.blocks, SetNowBlock)

}

//已知创世区块内容，从创世区块开始向下运算

//func (b *blockmain) SetBlockHash()
//SetBlockHash()用父哈希等算出来当前哈希

//func SetNowBlock(date string, LastBlockHash []byte) *blockmain
//SetNowBlock()将交易内容（date）和LastBlockHash写入当前的区块。

func main() {
	var bmc BlockMainChain
	firstblock()       //创造创世区块
	firstblockmainin() //创世区块进链

	//加一个循环反复创造下一个区块
	for j := 1; j <= 9; j++ {
		time.Sleep(3 * time.Second)

		for _, block := range &bmc.blocks {
			fmt.Printf("Prev. hash: %x\n", blockmain.LastBlockHash)
			fmt.Printf("Data: %s\n", blockmain.date)
			fmt.Printf("Hash: %x\n", blockmain.BlockHash)
			fmt.Println()
		}
		fmt.Println(j)
	}

}

//跑不动，不知道为什么。
