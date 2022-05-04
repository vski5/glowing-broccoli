# glowing-broccoli
go的区块链学习

# BTC-like
## 分两个struct分别实现区块和链。
### 区块
```go
type blockmain struct {

 whattime int64 //时间戳

 LastBlockHash []byte //父哈希

 BlockHash     []byte //当前哈希

 date          []byte //交易内容

}
```
## 链
```go
type BlockMainChain struct {

 blocks []*blockmain

}
```

区块链的链，是区块这个struct类型的指针的数组。

**类型指针**，允许对这个指针类型的数据进行修改。传递数据使用指针，而无须拷贝数据。

区块链的链，是区块这个struct类型的指针的数组。
```go
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
```
这里的数组，会依次增加，每一个元素都来自于区块。

**git** init https://github.com/vski5/glowing-broccoli.git