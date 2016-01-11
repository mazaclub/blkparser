package blkparser

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestParseBlock(t *testing.T) {
	// Arbitrarily chosen CLAM block.
	rawBlock, err := hex.DecodeString("0200000082ed2e1c1e3c159e3e4b0ac7db0770b96d1c110ec757be84d7599bb800000000ddfaa9db5a654b14ee34c45ccb65b470aa40177164b6955d137286259a281e148712f5523721021d0acb2d0f0101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0b01640101062f503253482fffffffff010088526a74000000232103569988948d05ddf970d610bc52f0d47fb21ec307a35d3cbeba6d11accfcd3c6aac00000000")
	if err != nil {
		t.Error(err)
	}

	block, err := NewBlock(rawBlock)
	if err != nil {
		t.Error(err)
	}

	if bytes.Equal(rawBlock, block.Raw) != true {
		t.Errorf("For raw block, expected %x, got %x", rawBlock, block.Raw)
	}
	if block.Hash != "000000009de1a49f76441af8be14d43d6ed9cc7ce919babab317521468c60482" {
		t.Error("For block hash, expected 000000009de1a49f76441af8be14d43d6ed9cc7ce919babab317521468c60482, got", block.Hash)
	}
	if block.Version != 2 {
		t.Error("For block version, expected 2, got", block.Version)
	}
	if block.MerkleRoot != "141e289a258672135d95b664711740aa70b465cb5cc434ee144b655adba9fadd" {
		t.Error("For merkle root, expected 141e289a258672135d95b664711740aa70b465cb5cc434ee144b655adba9fadd, got", block.MerkleRoot)
	}
	if block.BlockTime != 1391792775 {
		t.Error("For block time, expected 1391792775, got", block.BlockTime)
	}
	if block.Bits != 0x1d022137 {
		t.Error("For block bits, expected 0x1d022137, got", block.Bits)
	}
	if block.Nonce != 254659338 {
		t.Error("For block nonce, expected 254659338, got", block.Nonce)
	}
	if block.Size != 187 {
		t.Error("For block size, expected 187, got", block.Size)
	}
	if block.TxCnt != 1 {
		t.Error("For block tx count, expected 1, got", block.TxCnt)
	}
}
