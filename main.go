package main

import (
	"fmt"
	"github.com/soblom/go-lotto/lotto"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Usage: go-lotto <path_to_json>\n")
		os.Exit(0)
	}

	game, err := lotto.NewLottoGame(args[0])
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	for i := range game {
		fmt.Printf("%d\n", game[i])
	}
}
