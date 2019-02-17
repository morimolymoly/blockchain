package block

import (
	"encoding/hex"
	"fmt"
)

const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// Blockchain ... blockchain
type Blockchain struct {
	tip []byte
	Db  *blockdb
}

// FindUnspentTransactions ... get unspent transaction(output is not refed by any inputs)
func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block, err := bci.Next()
		if err != nil {
			panic(err)
		}

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlockedWith(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTXs
}

// FindUTXO ... get unspent outputs
func (bc *Blockchain) FindUTXO(address string) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := bc.FindUnspentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

// MineBlock ... mines a new block with the provided transactions
func (bc *Blockchain) MineBlock(transactions []*Transaction) error {
	tip, err := bc.Db.GetTip()
	if err != nil {
		return err
	}
	newBlock := NewBlock(transactions, tip)

	err = bc.Db.PutBlock(newBlock)
	return err
}

// NewGenesisBlock ... generate FIRTST block in our blockchain
func NewGenesisBlock(coinbase *Transaction) *Block {
	fmt.Printf("Hmm... There is no block in blockchain!\nLet's create Genesis Block!\n")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// NewBlockchain ... create blockchain with genesis block
func NewBlockchain(address string) (*Blockchain, error) {
	db, err := getNewBlockDB()
	if err != nil {
		return nil, err
	}

	null, err := db.CheckBlockchainIsNull()
	if err != nil {
		return nil, err
	}

	var tip []byte
	if null {
		db.CreateBlockchainBucket()
		cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)
		err := db.PutBlock(genesis)
		if err != nil {
			return nil, err
		}
		tip = genesis.Hash
	} else {
		h, err := db.GetTip()
		if err != nil {
			return nil, err
		}
		tip = h
	}
	return &Blockchain{
		tip: tip,
		Db:  db,
	}, nil
}
