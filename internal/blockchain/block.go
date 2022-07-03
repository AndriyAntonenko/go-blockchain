package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/AndriyAntonenko/blockchain/internal/utils"
)

type Block struct {
	Index     uint64
	PrevHash  string
	Timestamp *time.Time
	Data      string
	Proof     int64
}

func NewBlock(index uint64, timestamp time.Time, proof int64, prevHash string, data string) *Block {
	return &Block{
		Index:     index,
		Timestamp: &timestamp,
		Proof:     proof,
		PrevHash:  prevHash,
		Data:      data,
	}
}

func (b *Block) SerializeJSON() string {
	mapping := make(map[string]string)
	mapping["index"] = fmt.Sprint(b.Index)
	mapping["prevHash"] = b.PrevHash
	mapping["timestamp"] = b.Timestamp.String()
	mapping["data"] = b.Data
	mapping["proof"] = fmt.Sprint(b.Proof)
	mapping["hash"] = b.Hash()

	bytesJSON, err := json.Marshal(mapping)
	if err != nil {
		panic("cannot parse json " + err.Error())
	}
	return string(bytesJSON)
}

func (b *Block) Hash() string {
	stringToHash := strings.Join([]string{
		fmt.Sprint(b.Index),
		b.Timestamp.String(),
		b.PrevHash,
		b.Data,
		fmt.Sprint(b.Proof),
	}, ":")
	return utils.HashSha256ToHex(stringToHash)
}
