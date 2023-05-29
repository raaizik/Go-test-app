/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/notnil/chess"
	"github.com/notnil/chess/opening"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// findOpeningCmd represents the findOpening command
var findOpeningCmd = &cobra.Command{
	Use:   "findOpening",
	Short: "Returns the opening used in the received PGN",
	Long:  `Opens a path to a PGN files folder (PGN db), saved to PGNs folder in project's root dir, and returns the openings used according to the ECO.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No file name received as argument.")
			return
		}

		pgnDBFilename := args[0]

		// remove .pgn file extension if any
		if strings.Contains(pgnDBFilename, ".pgn") {
			pgnDBFilename = strings.ReplaceAll(pgnDBFilename, ".pgn", "")
		}
		pwd, _ := os.Getwd()
		dbPath := filepath.Join(pwd, "/PGNs/"+pgnDBFilename)
		exists, err := exists(dbPath)
		if !exists {
			log.Fatal("The db directory doesn't exist under the PGNs path!")
		} else if err != nil {
			log.Fatalf("See error: %v\n", err)
		}
		//path, err := os.Open(filepath.Join(pwd, "/PGNs/"+pgnDBFilename))

		items, _ := ioutil.ReadDir(dbPath)
		for _, item := range items {
			if item.IsDir() {
				log.Print("The database folder shouldn't contain folders")
				//fmt.Println("Bad structure. DB folder must contain only .pgn files")
				//return
				//subitems, _ := ioutil.ReadDir(item.Name())
				//for _, subitem := range subitems {
				//	if !subitem.IsDir() {
				//		// handle file there
				//		fmt.Println(item.Name() + "/" + subitem.Name())
				//	}
				//}
			} else {
				// handle file
				f, err := os.Open(filepath.Join(dbPath, item.Name()))
				if err != nil {
					panic(err)
				}

				// data leaks when calling defer inside the for loop (all files will remain open until function
				// returns)
				//defer f.Close()

				scanner := chess.NewScanner(f)

				game := chess.NewGame()
				//games := make([]*chess.Game, 0, 100)

				// scans first PGN in the file. If the file contains more than one PGN, it won't be included
				for scanner.Scan() {
					game = scanner.Next()
					//games = append(games, game)
				}

				//for _, game := range games {
				reader := strings.NewReader(game.String())
				pgn, err := chess.PGN(reader)
				if err != nil {
					// handle error
					log.Fatalf("Something went wrong, got error from %v", err)
				}

				g := chess.NewGame(pgn)

				// print opening name
				book := opening.NewBookECO()
				o := book.Find(g.Moves())
				fmt.Println(item.Name(), " --> ", o.Title())
				//}
				f.Close()
			}
		}
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
