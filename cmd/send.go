package cmd

import (
	"fmt"
	"os"
	"strconv"

	b "github.com/morimolymoly/blockchain/block"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&from, "from", "", "", "from")
	sendCmd.Flags().StringVarP(&to, "to", "", "", "to")
	sendCmd.Flags().StringVarP(&amount, "amount", "", "", "amount")
}

var sendCmd = &cobra.Command{
	Use:   "send coin",
	Short: "send coin",
	Long:  `send coin`,
	Run: func(cmd *cobra.Command, args []string) {
		bc, err := b.NewBlockchain(from)
		defer bc.Db.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		a, _ := strconv.Atoi(amount)
		tx := b.NewUTXOTransaction(from, to, a, bc)
		bc.MineBlock([]*b.Transaction{tx})
		fmt.Println("Success!")
	},
}
