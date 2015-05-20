package lotto

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const (
	json_valid   = "testdata/lotto_valid.json"
	json_invalid = "testdata/lotto_invalid.json"
)

var (
	valid_test_game = LottoGame{
		{[]int{1, 2, 3, 4, 5, 6}, 7},
		{[]int{1, 2, 3, 4, 5, 6}, 1},
		{[]int{1, 2, 3, 34, 38, 45}, 2}}
)

func TestReadNumbersValid(t *testing.T) {
	game, err := NewLottoGame(json_valid)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(game, valid_test_game) {
		t.Errorf(
			fmt.Sprintf(
				"'%s' json did not parse to LottoGame as expected", json_valid) +
				fmt.Sprintf("Actual content:\n%v\n", game) +
				fmt.Sprintf("Expected content:\n%v", valid_test_game))
	}
}

func TestReadNumbersInvalid(t *testing.T) {
	_, err := NewLottoGame(json_invalid)

	if err == nil {
		t.Fatal("Function failed to provoke any errors!")
	}

	uncaught := "\n"
	for _, v := range lottoError {
		if !strings.Contains(err.Error(), v) {
			uncaught += fmt.Sprintf("Error '%s' was not caught!\n", v)
		}
	}
	if uncaught != "" {
		t.Errorf(uncaught)
	}

}
