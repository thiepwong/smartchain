package core

import (
	"github.com/thiepwong/smartchain/core/params"
	"github.com/thiepwong/smartchain/core/types"
	"github.com/thiepwong/smartchain/smartdb/leveldb"
)

type BlockChain struct {
	chainConfig  params.ChainConfig
	db           leveldb.Database
	genesisBlock *types.Block
}

func NewBlockChain(db *leveldb.Database, chainConfig params.ChainConfig, genesis *types.Block) (*BlockChain, error) {
	bc := &BlockChain{
		chainConfig:  chainConfig,
		db:           *db,
		genesisBlock: genesis,
	}

	bc.addBlock(genesis)

	return bc, nil
}

func (bc *BlockChain) addBlock(block *types.Block) {

	// Validate the block
	bc.db.Save("so1", types.Bytes(block))
}

// func (bc *BlockChain) GetBlockByNumber(number uint64) *types.Block {

// 	hash := rawdb.ReadCanonicalHash(bc.db, number)
// 	if hash == (common.Hash{}) {
// 		return nil
// 	}
// 	return bc.GetBlock(hash, number)
// }

func (bc *BlockChain) PullChain() ([]byte, error) {
	return bc.db.Read("so1")
}

func GetLocalChain() (*BlockChain, error) {
	db := &leveldb.Database{}
	db.SetName("local-data")
	var err error
	db, err = db.Open()
	if err != nil {
		return nil, err
	}
	bc := &BlockChain{db: *db}
	return bc, nil
}

func GetlastBlock(bc *BlockChain) ([]byte, error) {
	b, e := bc.db.Read("so1")
	return b, e
}
