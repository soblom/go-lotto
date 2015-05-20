package lotto

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	in_num       = []int{6, 2, 1, 4, 5, 3}
	in_super_num = 7
)

func TestNewLottoRowNoErrorsOnValidInput(t *testing.T) {
	_, err := NewLottoRow(in_num, in_super_num)

	if err != nil {
		t.Errorf("Creation of valid lotto row (%d,%d) must not create error",
			in_num, in_super_num)
	}
}

func TestNewLottoRowSortsNumbers(t *testing.T) {
	in_num_sorted := []int{1, 2, 3, 4, 5, 6}

	row, _ := NewLottoRow(in_num, in_super_num)

	if !reflect.DeepEqual(row.Numbers, in_num_sorted) {
		t.Errorf("Row (%d) does not match sorted input row (%d)",
			row.Numbers, in_num_sorted)
	}
}

func TestNewLottoRowAllowsNoDuplicates(t *testing.T) {
	in_num_duplicate := []int{6, 2, 1, 3, 5, 3}

	_, err := NewLottoRow(in_num_duplicate, in_super_num)
	if err == nil {
		t.Errorf("'NewLottoRow' incorrectly accpeted input slice %d with duplicates!",
			in_num_duplicate)
	}
}

func TestSuperNumberOOB(t *testing.T) {
	wrong_sn := []int{10, 0, -3}
	error_strings := []string{}

	for n := range wrong_sn {
		_, err := NewLottoRow(in_num, wrong_sn[n])
		if err == nil {
			error_strings = append(error_strings,
				fmt.Sprintf("NewLottoRow allowed illegal "+
					"value %d for super_number\n", wrong_sn[n]))
		}
	}

	if len(error_strings) > 0 {
		t.Errorf("\n%s", error_strings)
	}
}

func TestNewLottoRowEnforcesSixNumbers(t *testing.T) {
	in_num_too_short := []int{6, 2, 1, 4, 5}
	in_num_too_long := []int{6, 2, 1, 4, 5, 3, 9}

	_, err := NewLottoRow(in_num_too_short, in_super_num)
	if err == nil {
		t.Errorf("NewLottoRow must check for presence of exactly 6 numbers "+
			"in lotto row. %d has %d elements",
			in_num_too_short, len(in_num_too_short))
	}

	_, err = NewLottoRow(in_num_too_long, in_super_num)
	if err == nil {
		t.Errorf("NewLottoRow must check for presence of exactly 6 numbers "+
			"in lotto row. %d has %d elements",
			in_num_too_long, len(in_num_too_long))
	}
}
