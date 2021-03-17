package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash

	h := sha256.New()
	h.Write([]byte(record))

	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock Block, data string) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index + 1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
