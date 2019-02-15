package cmd

import (
	"fmt"
	"os"
	"strconv"

	b "github.com/morimolymoly/blockchain/block"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(printchainCmd)
}

var printchainCmd = &cobra.Command{
	Use:   "printchain",
	Short: "print blockchain",
	Long:  `print blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		bc, err := b.NewBlockchain()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		i := bc.Iterator()
		for {
			block, err := i.Next()
			if err != nil {
				panic(err)
			}
			pow := b.NewProofOfWork(block)
			fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
			if i.Final {
				break
			}
		}
	},
}
