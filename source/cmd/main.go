package main

import (
	"github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal/cli"
	"github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal/storage/binance"
	"github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal/fetching"
	"github.com/spf13/cobra"
)

func main() {
	repo := binance.NewRepository()
	
	services := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "crypto-cli"}
	rootCmd.AddCommand(cli.InitCmd(services))
	rootCmd.Execute()
}