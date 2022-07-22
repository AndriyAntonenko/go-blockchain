package blockchain

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/AndriyAntonenko/blockchain/internal/utils"
)

const HASH_LEADING_ZEROS = 4

type Blockchain struct {
	chain []*Block
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{
		chain: []*Block{},
	}

	// add genesis block
	blockchain.CreateBlock(1, "0", "Genesis block!!!")
	return blockchain
}

func (b *Blockchain) CreateBlock(proof int64, prevHash string, data string) *Block {
	block := NewBlock(
		uint64(len(b.chain)+1),
		time.Now(),
		proof,
		prevHash,
		data,
	)

	b.chain = append(b.chain, block)
	return block
}

func (b *Blockchain) GetLastBlock() *Block {
	return b.chain[len(b.chain)-1]
}

func (b *Blockchain) ProofOfWork(prevProof int64) int64 {
	var newProof int64 = 1

	for {
		hash := b.puzzle(prevProof, newProof)
		if b.verifyPuzzle(hash) {
			break
		}
		newProof++
	}

	return newProof
}

func (b *Blockchain) puzzle(prevProof int64, newProof int64) string {
	solving := int64(math.Pow(float64(newProof), 2) - math.Pow(float64(prevProof), 2))
	return utils.HashSha256ToHex(fmt.Sprint(solving))
}

func (b *Blockchain) verifyPuzzle(puzzleHash string) bool {
	return puzzleHash[:HASH_LEADING_ZEROS] == strings.Repeat("0", HASH_LEADING_ZEROS)
}

func (b *Blockchain) isChainValid() bool {
	prevBlock := b.chain[0]
	block_index := 1

	for block_index < len(b.chain) {
		block := b.chain[block_index]
		if block.PrevHash != block.Hash() {
			return false
		}

		prevProof := prevBlock.Proof
		proof := block.Proof
		if !b.verifyPuzzle(b.puzzle(prevProof, proof)) {
			return false
		}

		prevBlock = block
		block_index++
	}

	return true
}

func (b *Blockchain) MineBlock(data string) *Block {
	prevBlock := b.GetLastBlock()
	prevProof := prevBlock.Proof
	proof := b.ProofOfWork(prevProof)
	block := b.CreateBlock(proof, prevBlock.Hash(), data)

	return block
}
