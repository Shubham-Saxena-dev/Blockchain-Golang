package blocks

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
)

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

const DIFFICULTY = 12

func (p *ProofOfWork) Run() (int, []byte) {

	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := p.InitData(nonce)
		hash := sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(p.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (p *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := p.InitData(p.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(p.Target) == -1
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-DIFFICULTY))

	pow := &ProofOfWork{b, target}

	return pow
}

func (p *ProofOfWork) InitData(nounce int) []byte {

	data := bytes.Join([][]byte{p.Block.PreviousHash, p.Block.Data,
		ToHex(int64(nounce)),
		ToHex(int64(DIFFICULTY))}, []byte{})
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
