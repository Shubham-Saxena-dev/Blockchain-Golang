package blocks

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

func InitBlock() *Block {
	return CreateBlock("First Block", []byte{})
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.Hash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	newBlock := &Block{
		Hash:         []byte{},
		Data:         []byte(data),
		PreviousHash: prevHash,
	}

	newBlock.DeriveHash()
	return newBlock
}
