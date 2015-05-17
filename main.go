package main

import (
        "fmt"
        "github.com/soblom/golotto/lotto"
)

func main() {
        fmt.Printf("Ich bin das Lotto-Programm\n")
        row, _ := lotto.NewLottoRow([6]int{1,2,3,4,5,6},7)
        fmt.Printf("%d, Superzahl: %d\n",row.Numbers,row.SuperNumber)
}
