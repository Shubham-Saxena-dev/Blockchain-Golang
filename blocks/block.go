package blocks

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce        int
}

func InitBlock() *Block {
	return CreateBlock("First Block", []byte{})
}

func CreateBlock(data string, prevHash []byte) *Block {
	newBlock := &Block{
		Hash:         []byte{},
		Data:         []byte(data),
		PreviousHash: prevHash,
		Nonce: 0,
	}

	pow := NewProof(newBlock)
	nonce, hash := pow.Run()

	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce
	return newBlock
}
