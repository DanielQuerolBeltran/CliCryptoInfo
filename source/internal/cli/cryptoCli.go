package cli

import (
	"log"

	app "github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal/fetching"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitCmd(service app.Service) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "cryptos",
		Short: "Print data about cryptos",
		Run:   runFn(service),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "Symbol of the crypto")

	return beersCmd
}

func runFn(service app.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetString(idFlag)
		if id != "" {
			crypto, err := service.GetCryptoById(id)
			if err != nil {
				log.Fatal(err)
			}
			crypto.Print()
			return
		}

		cryptos, err := service.GetAllCryptos()
		if err != nil {
			log.Fatal(err)
		}
		for _, crypto := range cryptos {
			crypto.Print()
		}
	}
}