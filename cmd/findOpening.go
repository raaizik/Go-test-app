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
	"os"
	"path/filepath"
	"strings"
)

// findOpeningCmd represents the findOpening command
var findOpeningCmd = &cobra.Command{
	Use:   "findOpening",
	Short: "Returns the opening used in the received PGN",
	Long:  `Opens a PGN file or a folder that is database of PGN files, saved to PGNs folder in project's root dir, and returns the opening used according to the ECO.`,
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
		//path, err := os.Open(filepath.Join(pwd, "/PGNs/"+pgnDBFilename))

		items, _ := ioutil.ReadDir(dbPath)
		for _, item := range items {
			if item.IsDir() {
				fmt.Println("Bad structure. DB folder must contain only .pgn files")
				//subitems, _ := ioutil.ReadDir(item.Name())
				//for _, subitem := range subitems {
				//	if !subitem.IsDir() {
				//		// handle file there
				//		fmt.Println(item.Name() + "/" + subitem.Name())
				//	}
				//}
			} else {
				// handle file there
				//fmt.Println(item.Name())
				f, err := os.Open(filepath.Join(dbPath, item.Name()))
				//fmt.Printf("File Name: %s\n", info.Name())
				//return nil
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
					fmt.Printf("Something went wrong, got error from %v", err)
				}

				g := chess.NewGame(pgn)

				// print opening name
				book := opening.NewBookECO()
				o := book.Find(g.Moves())
				fmt.Println(item.Name(), " - ", o.Title())
				//}
				f.Close()
			}
		}

		//filepath.Walk(dbPath, func(path string, info os.FileInfo, err error) error {
		//	if err != nil {
		//		log.Fatalf(err.Error())
		//	}
		//	f, err := os.Open(filepath.Join(dbPath, info.Name()))
		//	//fmt.Printf("File Name: %s\n", info.Name())
		//	//return nil
		//	if err != nil {
		//		panic(err)
		//	}
		//	defer f.Close()
		//
		//	scanner := chess.NewScanner(f)
		//	// todo: this grabs the first PGN from the db file.
		//	// can be more complex to have used choose which game from the db file to display
		//	// OR print opening for each one of the PGNs
		//	game := chess.NewGame()
		//	games := make([]*chess.Game, 0, 100)
		//
		//	for scanner.Scan() {
		//		game = scanner.Next()
		//		games = append(games, game)
		//	}
		//	// todo: still grabs only the first PGN in the db
		//	fmt.Println(len(games))
		//	for _, game := range games {
		//		reader := strings.NewReader(game.String())
		//		pgn, err := chess.PGN(reader)
		//		if err != nil {
		//			// handle error
		//			fmt.Printf("Something went wrong, got error from %v", err)
		//			return err
		//		}
		//
		//		g := chess.NewGame(pgn)
		//
		//		// print opening name
		//		book := opening.NewBookECO()
		//		o := book.Find(g.Moves())
		//		fmt.Println(o.Title())
		//	}
		//	return nil
		//})

		//f, err := os.Open("/Users/il015040/workspace/Go-test-app/PGNs/" + pgnDBFilename + "/*.pgn")

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
