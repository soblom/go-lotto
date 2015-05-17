package lotto

import (
        "testing"
_        "fmt"
)

func TestLottoRow(t *testing.T) {
        in_num := [6]int{6,2,1,4,5,3}
        in_num_sorted := [6]int{1,2,3,4,5,6}
        in_super_num := 7

        row, err := NewLottoRow(in_num, in_super_num)

        if err != nil {
                t.Errorf("Creation of valid lotto row (%d,%d) must not create error",
                        in_num, in_super_num)
        }

        if row.Numbers != in_num_sorted {
                t.Errorf("Row (%d) does not match sorted input row (%d)",
                        row.Numbers, in_num_sorted)
        }
}
