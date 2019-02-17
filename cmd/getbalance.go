package cmd

import (
	"fmt"
	"os"

	b "github.com/morimolymoly/blockchain/block"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getBalanceCmd)
	getBalanceCmd.Flags().StringVarP(&address, "address", "", "", "address")
}

var getBalanceCmd = &cobra.Command{
	Use:   "getbalance",
	Short: "get balance",
	Long:  `get balance`,
	Run: func(cmd *cobra.Command, args []string) {
		bc, err := b.NewBlockchain(address)
		defer bc.Db.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		balance := 0
		UTXOs := bc.FindUTXO(address)

		for _, out := range UTXOs {
			balance += out.Value
		}
	},
}
