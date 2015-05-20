package lotto

import (
        "fmt"
        "sort"
)

var lottoError map[string]string

func init() {
	lottoError = make(map[string]string)
	lottoError["6_numbers"] = "Each row needs to have exactly 6 numbers"
	lottoError["super_0..9"] = "The super number must be between 1 and 9"
	lottoError["num_no_dup"] = "numbers must not contain any duplicates"
}

// A LottoRow consists of 6 numbers between 1 and 49 and a `super number` between 0 and 9
type LottoRow struct {
	Numbers     []int
	SuperNumber int `json:"super"`
}

type LottoGame []LottoRow

type intArray []int
func (s intArray) Len() int { return len(s) }
func (s intArray) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s intArray) Less(i, j int) bool { return s[i] < s[j] }

func EmptyRow() LottoRow {
	return LottoRow{[]int{}, 0}
}

func NewLottoRow(numbers []int, super_number int) (result LottoRow, err error) {

        if len(numbers) < 6 {
                return EmptyRow(),
                        fmt.Errorf("Please specify 6 numbers for each lottery row")
        }

        var sorted_numbers []int = numbers

        return LottoRow{sort.Sort(numbers), super_number}, nil
}
