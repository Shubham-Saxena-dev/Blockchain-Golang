package blocks

type BlockChain struct {
	Blockchain []*Block
}

func InitBlockChain() *BlockChain {
	return &BlockChain{
		Blockchain: []*Block{InitBlock()},
	}
}

func (chain *BlockChain) AddBlock(data string) {

	lastBlockInChain := chain.Blockchain[len(chain.Blockchain)-1]
	newBlock := CreateBlock(data, lastBlockInChain.Hash)
	chain.Blockchain = append(chain.Blockchain, newBlock)
}
