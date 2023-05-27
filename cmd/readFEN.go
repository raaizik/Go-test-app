/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
	"github.com/spf13/cobra"
	"image/color"
	"log"
	"os"
	"path/filepath"
)

// readFENCmd represents the readFEN command
var readFENCmd = &cobra.Command{
	Use:   "readFEN",
	Short: "FEN to unicode symbols chess board + image",
	Long: `This command receives a FEN string and prints out the unicode symbols of the position + a .svg file containing a visual image of the current position on a board. File saved under boards folder under the project's root directory.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No FEN string received as argument.")
			return
		}
		fen, err := chess.FEN(args[0])
		if err != nil {
			// handle error
			fmt.Printf("Bad arg %v, got error from %v", args[0], err)
			return
		}
		game := chess.NewGame(fen)
		// print outcome and game PGN
		fmt.Println(game.Position().Board().Draw())
		fmt.Printf("Game position: %s by %s.\n", game.Outcome(), game.Method())
		fmt.Println(game.String())

		// create file
		pwd, _ := os.Getwd()
		FenStr := game.FEN()
		f, err := os.CreateTemp(filepath.Join(pwd, "/boards/"), "game-*.svg")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// create board position
		pos := &chess.Position{}
		if err := pos.UnmarshalText([]byte(FenStr)); err != nil {
			log.Fatal(err)
		}

		// write board SVG to file
		yellow := color.RGBA{255, 255, 0, 1}
		mark := image.MarkSquares(yellow, chess.D2, chess.D4)
		if err := image.SVG(f, pos.Board(), mark); err != nil {
			log.Fatal(err)
		}

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
