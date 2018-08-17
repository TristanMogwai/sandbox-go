package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
)

type Block struct {
	index        int
	timestamp    uint16
	hash         string
	previousHash string
	data         string
	nonce        int
}

func (block *Block) calculateHash() string {
	hash := sha256.Sum256([]byte(block.blockStringToHash()))

	return string(hash[:len(hash)])
}

func (block *Block) blockStringToHash() string {
	var buffer bytes.Buffer
	buffer.WriteString(string(block.index) + string(block.timestamp) + block.previousHash + block.data + string(block.nonce))
	return buffer.String()
}

func (block *Block) mineBlock(difficulty int) Block {
	block.nonce = 0

	var stringBuilder strings.Builder
	for i := 0; i < difficulty; i++ {
		fmt.Fprintf(&stringBuilder, "%d", 0)
	}

	for block.hash != stringBuilder.String() {
		block.nonce++
		block.calculateHash()
	}
}
