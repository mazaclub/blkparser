package blkparser

import (
	"os"
	"fmt"
	"bytes"
	"errors"
)

type Blockchain struct {
	Path string
	Magic [4]byte
	CurrentFile *os.File
	CurrentId uint32
}

func NewBlockchain(path string, magic [4]byte) (blockchain *Blockchain, err error) {
	blockchain = new(Blockchain)
	blockchain.Path = path
	blockchain.Magic = magic
	blockchain.CurrentId = 0

	f, err := os.Open(blkfilename(path, 0))
	if err != nil {
        return
    }

    blockchain.CurrentFile = f
	return
}

func (blockchain *Blockchain) NextBlock() (block *Block, err error) {
	rawblock, err := blockchain.FetchNextBlock()
	if err != nil {
		newblkfile, err2 := os.Open(blkfilename(blockchain.Path, blockchain.CurrentId+1))
		if err2 != nil {
			blockchain.CurrentId++
			blockchain.CurrentFile.Close()
			blockchain.CurrentFile = newblkfile
			rawblock, err = blockchain.FetchNextBlock()
		}
	}

	block, err = NewBlock(rawblock)
	if err != nil {
		return
	}

	return
}

func (blockchain *Blockchain) FetchNextBlock() (rawblock []byte, err error) {
	buf := [4]byte{}
    _, err = blockchain.CurrentFile.Read(buf[:])
    if err != nil {
        return
    }

    if !bytes.Equal(buf[:], blockchain.Magic[:]) {
    	err = errors.New("Bad magic") 
    	return
    }

    _, err = blockchain.CurrentFile.Read(buf[:])
    if err != nil {
        return
    }

    blocksize := uint32(blksize(buf[:]))

    rawblock = make([]byte, blocksize)

    _, err = blockchain.CurrentFile.Read(rawblock[:])
    if err != nil {
        return
    }

    return

}

func blkfilename(path string, id uint32) string {
	return fmt.Sprintf("%s/blk%05d.dat", path, id)
}

func blksize(buf []byte) (size uint64) {
        for i:=0; i<len(buf); i++ {
                size |= (uint64(buf[i]) << uint(i*8))
        }
        return
}