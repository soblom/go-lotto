package lotto

import (
        "fmt"
)

type LottoRow struct {
        Numbers [6]int
        SuperNumber int
}

type LottoGame []LottoRow

func EmptyRow() LottoRow {
        return LottoRow{[6]int{},0}
}

func NewLottoRow(numbers [6]int, super_number int) (result LottoRow, err error) {

        if len(numbers) < 6 {
                return EmptyRow(),
                        fmt.Errorf("Please specify 6 numbers for each lottery row")
        }
        return LottoRow{[6]int{},0}, nil
}
