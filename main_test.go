package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("ファイル読み込み成功", func(t *testing.T) {
		want := [][]string{
			{"5+5", "10"},
			{"1+1", "2"},
			{"8+3", "11"},
			{"1+2", "3"},
			{"8+6", "14"},
			{"3+1", "4"},
			{"1+4", "5"},
			{"5+1", "6"},
			{"2+3", "5"},
			{"3+3", "6"},
			{"2+4", "6"},
			{"5+2", "7"},
		}
		got, err := readCSVFile("problems.csv")
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
