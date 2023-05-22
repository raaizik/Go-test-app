/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/notnil/chess"
	"github.com/spf13/cobra"
	"math/rand"
	_ "math/rand"
	"time"
)

// getRandGameCmd represents the getRandGame command
var getRandGameCmd = &cobra.Command{
	Use:   "getRandGame",
	Short: "Prints out a random game",
	Long: `Returns a random game. Prints a visual board with unicode chess symbols, the game result and its PGN. Example output:
					/*
			Output:

			 A B C D E F G H
			8- - - - - - - -
			7- - - - - - ♚ -
			6- - - - ♗ - - -
			5- - - - - - - -
			4- - - - - - - -
			3♔ - - - - - - -
			2- - - - - - - -
			1- - - - - - - -

			Game completed. 1/2-1/2 by InsufficientMaterial.

			1.Nc3 b6 2.a4 e6 3.d4 Bb7 ...
		*/`,
	Run: func(cmd *cobra.Command, args []string) {
		game := chess.NewGame()
		// generate moves until game is over
		rand.Seed(time.Now().UnixNano())
		for game.Outcome() == chess.NoOutcome {
			// select a random move
			moves := game.ValidMoves()
			move := moves[rand.Intn(len(moves))]
			game.Move(move)
		}
		// print outcome and game PGN
		fmt.Println(game.Position().Board().Draw())
		fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
		fmt.Println(game.String())

	},
}

func init() {
	rootCmd.AddCommand(getRandGameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getRandGameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getRandGameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
