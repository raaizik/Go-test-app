/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/notnil/chess"

	"github.com/spf13/cobra"
)

// readFENCmd represents the readFEN command
var readFENCmd = &cobra.Command{
	Use:   "readFEN",
	Short: "FEN to unicode symbols chess board",
	Long: `This command receives a FEN string and prints out the unicode symbols of the position. For example:
`,
	Run: func(cmd *cobra.Command, args []string) {
		fen, err := chess.FEN(args[0])
		fmt.Println(args[0])
		if err != nil {
			// handle error
			fmt.Printf("Bad arg %v, %v\n", args[0])
		}
		game := chess.NewGame(fen)
		// print outcome and game PGN
		fmt.Println(game.Position().Board().Draw())
		fmt.Printf("Game position: %s by %s.\n", game.Outcome(), game.Method())
		fmt.Println(game.String())
	},
}

func init() {
	rootCmd.AddCommand(readFENCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readFENCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readFENCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
