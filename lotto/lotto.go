package lotto

import (
        "fmt"
        "sort"
)

type LottoRow struct {
        Numbers [6]int
        SuperNumber int
}

type LottoGame []LottoRow

type intArray []int
func (s intArray) Len() int { return len(s) }
func (s intArray) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s intArray) Less(i, j int) bool { return s[i] < s[j] }

func EmptyRow() LottoRow {
        return LottoRow{[6]int{},0}
}

func NewLottoRow(numbers []int, super_number int) (result LottoRow, err error) {

        if len(numbers) < 6 {
                return EmptyRow(),
                        fmt.Errorf("Please specify 6 numbers for each lottery row")
        }

        var sorted_numbers []int = numbers

        return LottoRow{sort.Sort(numbers), super_number}, nil
}
