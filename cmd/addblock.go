package cmd

import (
	"fmt"
	"os"

	b "github.com/morimolymoly/blockchain/block"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addblockCmd)
	addblockCmd.Flags().StringVarP(&data, "data", "", "", "data")
}

var data string

var addblockCmd = &cobra.Command{
	Use:   "addblock",
	Short: "add block to blockchain",
	Long:  `add block to blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		bc, err := b.NewBlockchain()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = bc.AddBlock(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
