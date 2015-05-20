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

// `NewLottoRow` creates a `LottoRow` with `numbers` and `super_number` as values.
// The `numbers` slice needs to contain exactly six different integers and the
// `super_number` must be a value between `1` and `9`
// It returns a new `LottoRow` object in which the `numbers` will be sorted.
func NewLottoRow(numbers []int, super_number int) (result LottoRow, err error) {

	if len(numbers) != 6 {
		return EmptyRow(),
			fmt.Errorf(lottoError["6_numbers"])
	}

	if (super_number < 1) || (super_number > 9) {
		return EmptyRow(),
			fmt.Errorf(lottoError["super_0..9"])
	}

	sort.Sort(sort.IntSlice(numbers))

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] == numbers[i+1] {
			return EmptyRow(),
				fmt.Errorf(lottoError["num_no_dup"])
		}
	}

	return LottoRow{numbers, super_number}, nil
}

        var sorted_numbers []int = numbers
func NewLottoGame(path string) (LottoGame, error) {
	data, err := loadRawDatafromJSON(path)
	if err != nil {
		return nil, err
	}

	row_errors := make([]string, 0, len(data))
	game := make([]LottoRow, len(data), len(data))
	for i := range data {
		row, err := NewLottoRow(data[i].Numbers, data[i].SuperNumber)
		if err != nil {
			row_errors = append(row_errors, err.Error())
		} else {
			game[i] = row
		}
	}

	if len(row_errors) > 0 {
		return nil, fmt.Errorf(strings.Join(row_errors, "\n"))
	}

	return game, nil
}

func loadRawDatafromJSON(path string) (LottoGame, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file.\n%s", err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from file.\n%s", err)
	}

	var data LottoGame
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse JSON.\n%s", err)
	}

        return LottoRow{sort.Sort(numbers), super_number}, nil
}
