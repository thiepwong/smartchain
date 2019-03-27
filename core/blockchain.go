package core

import (
	"github.com/thiepwong/smartchain/core/params"
	"github.com/thiepwong/smartchain/core/types"
	"github.com/thiepwong/smartchain/smartdb/mongodb"
)

//BlockChain struct
type BlockChain struct {
	chainConfig  params.ChainConfig
	db           *mongodb.Database
	genesisBlock *types.Block
}

//NewBlockChain func to create a new blockchain
func NewBlockChain(db *mongodb.Database, chainConfig params.ChainConfig, genesis *types.Block) (*BlockChain, error) {
	bc := &BlockChain{
		chainConfig:  chainConfig,
		db:           db,
		genesisBlock: genesis,
	}

	bc.addBlock(genesis)

	return bc, nil
}

func (bc *BlockChain) addBlock(block *types.Block) error {

	// Validate the block
	return bc.db.Insert("mainchain", block)
}

// func (bc *BlockChain) GetBlockByNumber(number uint64) *types.Block {

// 	hash := rawdb.ReadCanonicalHash(bc.db, number)
// 	if hash == (common.Hash{}) {
// 		return nil
// 	}
// 	return bc.GetBlock(hash, number)
// }

// func (bc *BlockChain) PullChain() ([]byte, error) {
// 	return bc.db.Read("so1")
// }

// func GetLocalChain() (*BlockChain, error) {
// 	db := &mongodb.Database{}

// 	db, err = db.Open()
// 	if err != nil {
// 		return nil, err
// 	}
// 	bc := &BlockChain{db: *db}
// 	return bc, nil
// }

// func GetlastBlock(bc *BlockChain) ([]byte, error) {
// 	b, e := bc.db.Read("so1")
// 	return b, e
// }
