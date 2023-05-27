/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/notnil/chess"
	"github.com/notnil/chess/opening"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// findOpeningCmd represents the findOpening command
var findOpeningCmd = &cobra.Command{
	Use:   "findOpening",
	Short: "Returns the opening used in the received PGN",
	Long:  `Opens a PGN file name (w/o .pgn extension), saved to PGNs folder in project's root dir, of a game and returns the opening used according to the ECO.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No PGN file name received as argument.")
			return
		}

		pgnFilename := args[0]

		// remove .pgn file extension if any
		if strings.Contains(pgnFilename, ".pgn") {
			pgnFilename = strings.ReplaceAll(pgnFilename, ".pgn", "")
		}
		f, err := os.Open("/Users/il015040/workspace/Go-test-app/PGNs/" + pgnFilename + ".pgn")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		scanner := chess.NewScanner(f)
		game := chess.NewGame()
		for scanner.Scan() {
			game = scanner.Next()
		}

		pgn, err := chess.PGN(strings.NewReader(game.String()))
		if err != nil {
			// handle error
			fmt.Printf("Something went wrong, got error from %v", err)
			return
		}

		g := chess.NewGame(pgn)

		// print opening name
		book := opening.NewBookECO()
		o := book.Find(g.Moves())
		fmt.Println(o.Title())
	},
}

func init() {
	rootCmd.AddCommand(findOpeningCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findOpeningCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findOpeningCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
